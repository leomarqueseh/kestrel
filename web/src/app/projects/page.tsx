"use client";

import { useEffect, useState, FormEvent } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import { getToken, listProjects, createProject, Project } from "@/lib/api";
import Sidebar from "@/components/Sidebar";

export default function ProjectsPage() {
  const router = useRouter();
  const [ready, setReady] = useState(false);
  const [projects, setProjects] = useState<Project[]>([]);
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [creating, setCreating] = useState(false);

  useEffect(() => {
    if (!getToken()) {
      router.push("/login");
      return;
    }
    setReady(true);
    refresh();
  }, [router]);

  function refresh() {
    listProjects()
      .then((data) => setProjects(data ?? []))
      .catch((err) => setError(err.message));
  }

  async function handleCreate(e: FormEvent) {
    e.preventDefault();
    setError(null);
    setCreating(true);
    try {
      await createProject(name, description);
      setName("");
      setDescription("");
      refresh();
    } catch (err) {
      setError(err instanceof Error ? err.message : "Failed to create project");
    } finally {
      setCreating(false);
    }
  }

  if (!ready) return null;

  return (
    <div className="flex min-h-screen bg-slate-950 text-slate-100">
      <Sidebar />
      <main className="flex-1 p-8">
        <h1 className="mb-6 text-2xl font-semibold">Projects</h1>

        <form onSubmit={handleCreate} className="mb-8 flex flex-wrap gap-3 rounded-lg border border-slate-800 bg-slate-900 p-4">
          <input
            placeholder="Project name"
            required
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="min-w-[180px] flex-1 rounded border border-slate-700 bg-slate-950 px-3 py-2 text-sm"
          />
          <input
            placeholder="Description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="min-w-[220px] flex-[2] rounded border border-slate-700 bg-slate-950 px-3 py-2 text-sm"
          />
          <button
            type="submit"
            disabled={creating}
            className="rounded bg-emerald-600 px-4 py-2 text-sm font-medium text-white hover:bg-emerald-500 disabled:opacity-50"
          >
            {creating ? "Creating..." : "New project"}
          </button>
        </form>

        {error && <p className="mb-4 text-sm text-red-400">{error}</p>}

        <div className="space-y-2">
          {projects.map((p) => (
            <Link
              key={p.id}
              href={`/projects/${p.id}`}
              className="block rounded-lg border border-slate-800 bg-slate-900 p-4 hover:border-slate-700"
            >
              <p className="font-medium text-slate-100">{p.name}</p>
              <p className="text-sm text-slate-400">{p.description}</p>
            </Link>
          ))}
          {projects.length === 0 && <p className="text-slate-400">No projects yet.</p>}
        </div>
      </main>
    </div>
  );
}
