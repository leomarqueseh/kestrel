"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { getToken, listProjects, getDashboard, Project, DashboardSummary } from "@/lib/api";
import Sidebar from "@/components/Sidebar";

export default function DashboardPage() {
  const router = useRouter();
  const [ready, setReady] = useState(false);
  const [projects, setProjects] = useState<Project[]>([]);
  const [selected, setSelected] = useState<string>("");
  const [summary, setSummary] = useState<DashboardSummary | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!getToken()) {
      router.push("/login");
      return;
    }
    setReady(true);
    listProjects()
      .then((data) => {
        setProjects(data ?? []);
        if (data && data.length > 0) setSelected(data[0].id);
      })
      .catch((err) => setError(err.message));
  }, [router]);

  useEffect(() => {
    if (!selected) return;
    getDashboard(selected).then(setSummary).catch((err) => setError(err.message));
  }, [selected]);

  if (!ready) return null;

  return (
    <div className="flex min-h-screen bg-slate-950 text-slate-100">
      <Sidebar />
      <main className="flex-1 p-8">
        <div className="mb-6 flex items-center justify-between">
          <h1 className="text-2xl font-semibold">Dashboard</h1>
          {projects.length > 0 && (
            <select
              value={selected}
              onChange={(e) => setSelected(e.target.value)}
              className="rounded border border-slate-700 bg-slate-900 px-3 py-2 text-sm"
            >
              {projects.map((p) => (
                <option key={p.id} value={p.id}>{p.name}</option>
              ))}
            </select>
          )}
        </div>

        {error && <p className="mb-4 text-sm text-red-400">{error}</p>}

        {projects.length === 0 && (
          <p className="text-slate-400">
            No projects yet. Head to{" "}
            <a href="/projects" className="text-emerald-400 underline">Projects</a> to create one.
          </p>
        )}

        {summary && (
          <>
            <div className="mb-6 grid grid-cols-2 gap-4 md:grid-cols-4">
              <StatCard label="Targets" value={summary.total_targets} />
              <StatCard label="Scans" value={summary.total_scans} />
              <StatCard label="Confirmed findings" value={summary.confirmed_findings} />
              <StatCard label="Open findings" value={summary.open_findings} />
            </div>

            <div className="rounded-lg border border-slate-800 bg-slate-900 p-6">
              <h2 className="mb-4 text-sm font-medium text-slate-300">Findings by severity</h2>
              <div className="space-y-1">
                <SeverityRow label="Critical" value={summary.critical} dot="bg-red-500" />
                <SeverityRow label="High" value={summary.high} dot="bg-orange-500" />
                <SeverityRow label="Medium" value={summary.medium} dot="bg-amber-500" />
                <SeverityRow label="Low" value={summary.low} dot="bg-lime-500" />
                <SeverityRow label="Informational" value={summary.informational} dot="bg-slate-500" />
              </div>
            </div>
          </>
        )}
      </main>
    </div>
  );
}

function StatCard({ label, value }: { label: string; value: number }) {
  return (
    <div className="rounded-lg border border-slate-800 bg-slate-900 p-4">
      <p className="text-xs text-slate-400">{label}</p>
      <p className="mt-1 text-2xl font-semibold">{value}</p>
    </div>
  );
}

function SeverityRow({ label, value, dot }: { label: string; value: number; dot: string }) {
  return (
    <div className="flex items-center justify-between border-b border-slate-800 py-2 text-sm last:border-0">
      <span className="flex items-center gap-2 text-slate-300">
        <span className={`h-2 w-2 rounded-full ${dot}`} />
        {label}
      </span>
      <span className="text-slate-100">{value}</span>
    </div>
  );
}
