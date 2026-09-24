const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL ?? "http://localhost:8080";

export function getToken(): string | null {
  if (typeof window === "undefined") return null;
  return localStorage.getItem("kestrel_token");
}

export function setToken(token: string) {
  localStorage.setItem("kestrel_token", token);
}

export function clearToken() {
  localStorage.removeItem("kestrel_token");
}

export function getRole(): string | null {
  const token = getToken();
  if (!token) return null;
  try {
    const payload = JSON.parse(
      atob(token.split(".")[1].replace(/-/g, "+").replace(/_/g, "/"))
    );
    return payload.role ?? null;
  } catch {
    return null;
  }
}

export async function apiFetch<T>(path: string, options: RequestInit = {}): Promise<T> {
  const token = getToken();
  const headers = new Headers(options.headers);
  headers.set("Content-Type", "application/json");
  if (token) headers.set("Authorization", `Bearer ${token}`);

  const res = await fetch(`${API_BASE_URL}${path}`, { ...options, headers });

  if (!res.ok) {
    if (res.status === 401 && typeof window !== "undefined") {
      clearToken();
      window.location.href = "/login";
    }
    const body = await res.json().catch(() => ({}));
    throw new Error(body.error ?? `Request failed with status ${res.status}`);
  }
  if (res.status === 204) return undefined as T;
  return res.json();
}

export async function login(email: string, password: string) {
  const data = await apiFetch<{ access_token: string; refresh_token: string }>(
    "/api/v1/auth/login",
    { method: "POST", body: JSON.stringify({ email, password }) }
  );
  setToken(data.access_token);
  return data;
}

export interface Project {
  id: string; name: string; description: string; owner_id: string; created_at: string;
}
export interface Target {
  id: string; project_id: string; value: string; target_type: string;
  authorized: boolean; description: string; created_at: string;
}
export interface Scan {
  id: string; target_id: string; module: string; status: string;
  started_at?: string; finished_at?: string; error?: string; created_at: string;
}
export interface Finding {
  id: string; asset_id: string; title: string; description: string;
  severity: string; cvss?: number; status: string; recommendation: string; created_at: string;
}
export interface DashboardSummary {
  total_targets: number; total_scans: number;
  critical: number; high: number; medium: number; low: number; informational: number;
  confirmed_findings: number; open_findings: number; resolved_findings: number;
}

export const listProjects = () => apiFetch<Project[]>("/api/v1/projects");
export const createProject = (name: string, description: string) =>
  apiFetch<Project>("/api/v1/projects", { method: "POST", body: JSON.stringify({ name, description }) });

export const listTargets = (projectId: string) =>
  apiFetch<Target[]>(`/api/v1/projects/${projectId}/targets`);
export const createTarget = (projectId: string, value: string, target_type: string, description: string) =>
  apiFetch<Target>(`/api/v1/projects/${projectId}/targets`, {
    method: "POST",
    body: JSON.stringify({ value, target_type, description }),
  });
export const authorizeTarget = (targetId: string) =>
  apiFetch<Target>(`/api/v1/targets/${targetId}/authorize`, { method: "POST" });

export const runRecon = (targetId: string) =>
  apiFetch<Scan>(`/api/v1/targets/${targetId}/scans/recon`, { method: "POST" });
export const runEnumeration = (targetId: string) =>
  apiFetch<Scan>(`/api/v1/targets/${targetId}/scans/enumeration`, { method: "POST" });
export const listScans = (targetId: string) =>
  apiFetch<Scan[]>(`/api/v1/targets/${targetId}/scans`);

export const runAssessment = (targetId: string) =>
  apiFetch<Finding[]>(`/api/v1/targets/${targetId}/assessment`, { method: "POST" });
export const listFindingsByTarget = (targetId: string) =>
  apiFetch<Finding[]>(`/api/v1/targets/${targetId}/findings`);

export const startValidation = (findingId: string) =>
  apiFetch<Finding>(`/api/v1/findings/${findingId}/start-validation`, { method: "POST" });
export const confirmFinding = (findingId: string, request: string, response: string, notes: string) =>
  apiFetch<Finding>(`/api/v1/findings/${findingId}/confirm`, {
    method: "POST",
    body: JSON.stringify({ request, response, notes }),
  });
export const rejectFinding = (findingId: string, reason: string) =>
  apiFetch<Finding>(`/api/v1/findings/${findingId}/reject`, {
    method: "POST",
    body: JSON.stringify({ reason }),
  });

export const getDashboard = (projectId: string) =>
  apiFetch<DashboardSummary>(`/api/v1/projects/${projectId}/dashboard`);

export async function downloadReport(projectId: string, format: "html" | "json") {
  const token = getToken();
  const res = await fetch(`${API_BASE_URL}/api/v1/projects/${projectId}/report?format=${format}`, {
    headers: token ? { Authorization: `Bearer ${token}` } : {},
  });
  if (!res.ok) throw new Error(`Report generation failed (${res.status})`);
  const blob = await res.blob();
  window.open(URL.createObjectURL(blob), "_blank");
}
