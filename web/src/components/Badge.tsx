const severityColors: Record<string, string> = {
  critical: "bg-red-950 text-red-400 border-red-800",
  high: "bg-orange-950 text-orange-400 border-orange-800",
  medium: "bg-amber-950 text-amber-400 border-amber-800",
  low: "bg-lime-950 text-lime-400 border-lime-800",
  informational: "bg-slate-800 text-slate-400 border-slate-700",
};

const statusColors: Record<string, string> = {
  detected: "bg-slate-800 text-slate-400 border-slate-700",
  needs_validation: "bg-blue-950 text-blue-400 border-blue-800",
  confirmed: "bg-emerald-950 text-emerald-400 border-emerald-800",
  false_positive: "bg-slate-800 text-slate-500 border-slate-700 line-through",
  pending: "bg-slate-800 text-slate-400 border-slate-700",
  running: "bg-blue-950 text-blue-400 border-blue-800",
  completed: "bg-emerald-950 text-emerald-400 border-emerald-800",
  failed: "bg-red-950 text-red-400 border-red-800",
};

export function SeverityBadge({ value }: { value: string }) {
  return (
    <span className={`inline-block rounded border px-2 py-0.5 text-xs ${severityColors[value] ?? severityColors.informational}`}>
      {value}
    </span>
  );
}

export function StatusBadge({ value }: { value: string }) {
  return (
    <span className={`inline-block rounded border px-2 py-0.5 text-xs ${statusColors[value] ?? statusColors.pending}`}>
      {value}
    </span>
  );
}
