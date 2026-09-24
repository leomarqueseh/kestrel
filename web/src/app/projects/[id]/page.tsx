"use client";

import { useEffect, useState, FormEvent } from "react";
import { useParams, useRouter } from "next/navigation";
import {
  getToken, getRole,
  listTargets, createTarget, authorizeTarget,
  runRecon, runEnumeration, listScans,
  runAssessment, listFindingsByTarget,
  startValidation, confirmFinding, rejectFinding,
  downloadReport,
  Target, Scan, Finding,
} from "@/lib/api";
import Sidebar from "@/components/Sidebar";
import { SeverityBadge, StatusBadge } from "@/components/Badge";

export default function ProjectDetailPage() {
  const params = useParams();
  const projectId = params.id as string;
  const router = useRouter();
  const [ready, setReady] = useState(false);
  const role = getRole();

  const [targets, setTargets] = useState<Target[]>([]);
  const [selectedTarget, setSelectedTarget] = useState<string>("");
  const [scans, setScans] = useState<Scan[]>([]);
  const [findings, setFindings] = useState<Finding[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [busy, setBusy] = useState<string | null>(null);

  const [newValue, setNewValue] = useState("");
  const [newType, setNewType] = useState("domain");
  const [newDescription, setNewDescription] = useState("");

  useEffect(() => {
    if (!getToken()) {
      router.push("/login");
      return;
    }
    setReady(true);
    refreshTargets();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [router]);

  useEffect(() => {
    if (!selectedTarget) return;
    refreshScansAndFindings();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [selectedTarget]);

  function refreshTargets() {
    listTargets(projectId)
      .then((data) => {
        const list = data ?? [];
        setTargets(list);
        if (list.length > 0 && !selectedTarget) setSelectedTarget(list[0].id);
      })
      .catch((err) => setError(err.message));
  }

  // Go encodes an empty/nil slice as JSON `null`, not `[]` — every list
  // response here falls back to an empty array so `.map()` never blows up
  // on a target/project with no scans or findings yet.
  function refreshScansAndFindings() {
    listScans(selectedTarget)
      .then((data) => setScans(data ?? []))
      .catch((err) => setError(err.message));
    listFindingsByTarget(selectedTarget)
      .then((data) => setFindings(data ?? []))
      .catch((err) => setError(err.message));
  }

  async function handleCreateTarget(e: FormEvent) {
    e.preventDefault();
    setError(null);
    try {
      await createTarget(projectId, newValue, newType, newDescription);
      setNewValue("");
      setNewDescription("");
      refreshTargets();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to create target");
    }
  }

  async function handleAuthorize(targetId: string) {
    setError(null);
    setBusy(targetId);
    try {
      await authorizeTarget(targetId);
      refreshTargets();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to authorize target");
    } finally {
      setBusy(null);
    }
  }

  async function handleScan(kind: "recon" | "enumeration") {
    if (!selectedTarget) return;
    setError(null);
    setBusy(kind);
    try {
      if (kind === "recon") await runRecon(selectedTarget);
      else await runEnumeration(selectedTarget);
      refreshScansAndFindings();
    } catch (err) {
      setError(err instanceof Error ? err.message : `Failed to run ${kind}`);
    } finally {
      setBusy(null);
    }
  }

  async function handleAssess() {
    if (!selectedTarget) return;
    setError(null);
    setBusy("assessment");
    try {
      await runAssessment(selectedTarget);
      refreshScansAndFindings();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Assessment failed");
    } finally {
      setBusy(null);
    }
  }

  async function handleStartValidation(findingId: string) {
    setBusy(findingId);
    try {
      await startValidation(findingId);
      refreshScansAndFindings();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to start validation");
    } finally {
      setBusy(null);
    }
  }

  async function handleConfirm(findingId: string) {
    const request = window.prompt("Request sent:", "GET / HTTP/1.1") ?? "";
    const response = window.prompt("Response received:", "") ?? "";
    const notes = window.prompt("Notes:", "") ?? "";
    setBusy(findingId);
    try {
      await confirmFinding(findingId, request, response, notes);
      refreshScansAndFindings();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to confirm finding");
    } finally {
      setBusy(null);
    }
  }

  async function handleReject(findingId: string) {
    const reason = window.prompt("Reason for rejecting:", "") ?? "";
    setBusy(findingId);
    try {
      await rejectFinding(findingId, reason);
      refreshScansAndFindings();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to reject finding");
    } finally {
      setBusy(null);
    }
  }

  if (!ready) return null;

  const canAuthorize = role === "admin";
  const canOperate = role === "admin" || role === "analyst";

  return (
    <div className="flex min-h-screen bg-slate-950 text-slate-100">
      <Sidebar />
      <main className="flex-1 space-y-8 p-8">
        <div className="flex items-center justify-between">
          <h1 className="text-2xl font-semibold">Project targets</h1>
          <div className="flex gap-2">
            <button
              onClick={() => downloadReport(projectId, "html").catch((e) => setError(e.message))}
              className="rounded border border-slate-700 px-3 py-2 text-sm hover:bg-slate-800"
            >
              View report (HTML)
            </button>
            <button
              onClick={() => downloadReport(projectId, "json").catch((e) => setError(e.message))}
              className="rounded border border-slate-700 px-3 py-2 text-sm hover:bg-slate-800"
            >
              Export JSON
            </button>
          </div>
        </div>

        {error && <p className="text-sm text-red-400">{error}</p>}

        {canOperate && (
          <form onSubmit={handleCreateTarget} className="flex flex-wrap items-end gap-3 rounded-lg border border-slate-800 bg-slate-900 p-4">
            <div className="min-w-[160px] flex-1">
              <label className="mb-1 block text-xs text-slate-400">Value</label>
              <input
                required
                value={newValue}
                onChange={(e) => setNewValue(e.target.value)}
                placeholder="example.com"
                className="w-full rounded border border-slate-700 bg-slate-950 px-3 py-2 text-sm"
              />
            </div>
            <div>
              <label className="mb-1 block text-xs text-slate-400">Type</label>
              <select
                value={newType}
                onChange={(e) => setNewType(e.target.value)}
                className="rounded border border-slate-700 bg-slate-950 px-3 py-2 text-sm"
              >
                <option value="domain">domain</option>
                <option value="ip">ip</option>
                <option value="url">url</option>
                <option value="cidr">cidr</option>
              </select>
            </div>
            <div className="min-w-[200px] flex-[2]">
              <label className="mb-1 block text-xs text-slate-400">Description</label>
              <input
                value={newDescription}
                onChange={(e) => setNewDescription(e.target.value)}
                className="w-full rounded border border-slate-700 bg-slate-950 px-3 py-2 text-sm"
              />
            </div>
            <button
              type="submit"
              className="rounded bg-emerald-600 px-4 py-2 text-sm font-medium text-white hover:bg-emerald-500"
            >
              Add target
            </button>
          </form>
        )}

        <div className="rounded-lg border border-slate-800 bg-slate-900">
          <table className="w-full text-sm">
            <thead>
              <tr className="border-b border-slate-800 text-left text-xs text-slate-400">
                <th className="p-3">Value</th>
                <th className="p-3">Type</th>
                <th className="p-3">Authorized</th>
                <th className="p-3"></th>
              </tr>
            </thead>
            <tbody>
              {targets.map((t) => (
                <tr
                  key={t.id}
                  onClick={() => setSelectedTarget(t.id)}
                  className={`cursor-pointer border-b border-slate-800 last:border-0 hover:bg-slate-800/50 ${
                    selectedTarget === t.id ? "bg-slate-800/70" : ""
                  }`}
                >
                  <td className="p-3">{t.value}</td>
                  <td className="p-3 text-slate-400">{t.target_type}</td>
                  <td className="p-3">
                    {t.authorized ? (
                      <span className="text-emerald-400">Yes</span>
                    ) : (
                      <span className="text-red-400">No</span>
                    )}
                  </td>
                  <td className="p-3 text-right">
                    {!t.authorized && canAuthorize && (
                      <button
                        onClick={(e) => {
                          e.stopPropagation();
                          handleAuthorize(t.id);
                        }}
                        disabled={busy === t.id}
                        className="rounded border border-emerald-700 px-2 py-1 text-xs text-emerald-400 hover:bg-emerald-950 disabled:opacity-50"
                      >
                        {busy === t.id ? "Authorizing..." : "Authorize"}
                      </button>
                    )}
                  </td>
                </tr>
              ))}
              {targets.length === 0 && (
                <tr>
                  <td colSpan={4} className="p-4 text-center text-slate-400">
                    No targets yet.
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>

        {selectedTarget && (
          <>
            <div>
              <div className="mb-3 flex items-center justify-between">
                <h2 className="text-lg font-medium">Scans</h2>
                {canOperate && (
                  <div className="flex gap-2">
                    <button
                      onClick={() => handleScan("recon")}
                      disabled={busy === "recon"}
                      className="rounded border border-slate-700 px-3 py-1.5 text-xs hover:bg-slate-800 disabled:opacity-50"
                    >
                      {busy === "recon" ? "Running..." : "Run recon"}
                    </button>
                    <button
                      onClick={() => handleScan("enumeration")}
                      disabled={busy === "enumeration"}
                      className="rounded border border-slate-700 px-3 py-1.5 text-xs hover:bg-slate-800 disabled:opacity-50"
                    >
                      {busy === "enumeration" ? "Running..." : "Run enumeration"}
                    </button>
                    <button
                      onClick={handleAssess}
                      disabled={busy === "assessment"}
                      className="rounded bg-emerald-600 px-3 py-1.5 text-xs font-medium text-white hover:bg-emerald-500 disabled:opacity-50"
                    >
                      {busy === "assessment" ? "Assessing..." : "Assess vulnerabilities"}
                    </button>
                  </div>
                )}
              </div>
              <div className="rounded-lg border border-slate-800 bg-slate-900">
                <table className="w-full text-sm">
                  <thead>
                    <tr className="border-b border-slate-800 text-left text-xs text-slate-400">
                      <th className="p-3">Module</th>
                      <th className="p-3">Status</th>
                      <th className="p-3">Created</th>
                    </tr>
                  </thead>
                  <tbody>
                    {scans.map((s) => (
                      <tr key={s.id} className="border-b border-slate-800 last:border-0">
                        <td className="p-3">{s.module}</td>
                        <td className="p-3"><StatusBadge value={s.status} /></td>
                        <td className="p-3 text-slate-400">{new Date(s.created_at).toLocaleString()}</td>
                      </tr>
                    ))}
                    {scans.length === 0 && (
                      <tr><td colSpan={3} className="p-4 text-center text-slate-400">No scans yet.</td></tr>
                    )}
                  </tbody>
                </table>
              </div>
            </div>

            <div>
              <h2 className="mb-3 text-lg font-medium">Findings</h2>
              <div className="space-y-3">
                {findings.map((f) => (
                  <div key={f.id} className="rounded-lg border border-slate-800 bg-slate-900 p-4">
                    <div className="mb-2 flex items-center justify-between">
                      <p className="font-medium">{f.title}</p>
                      <div className="flex gap-2">
                        <SeverityBadge value={f.severity} />
                        <StatusBadge value={f.status} />
                      </div>
                    </div>
                    <p className="mb-2 text-sm text-slate-400">{f.description}</p>
                    {f.cvss ? <p className="mb-2 text-xs text-slate-500">CVSS: {f.cvss}</p> : null}

                    {canOperate && (
                      <div className="flex gap-2">
                        {f.status === "detected" && (
                          <button
                            onClick={() => handleStartValidation(f.id)}
                            disabled={busy === f.id}
                            className="rounded border border-blue-800 px-2 py-1 text-xs text-blue-400 hover:bg-blue-950 disabled:opacity-50"
                          >
                            Start validation
                          </button>
                        )}
                        {f.status === "needs_validation" && (
                          <>
                            <button
                              onClick={() => handleConfirm(f.id)}
                              disabled={busy === f.id}
                              className="rounded border border-emerald-800 px-2 py-1 text-xs text-emerald-400 hover:bg-emerald-950 disabled:opacity-50"
                            >
                              Confirm
                            </button>
                            <button
                              onClick={() => handleReject(f.id)}
                              disabled={busy === f.id}
                              className="rounded border border-slate-700 px-2 py-1 text-xs text-slate-400 hover:bg-slate-800 disabled:opacity-50"
                            >
                              Reject
                            </button>
                          </>
                        )}
                      </div>
                    )}
                  </div>
                ))}
                {findings.length === 0 && <p className="text-slate-400">No findings yet.</p>}
              </div>
            </div>
          </>
        )}
      </main>
    </div>
  );
}
