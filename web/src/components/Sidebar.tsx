"use client";

import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { clearToken } from "@/lib/api";

const links = [
  { href: "/dashboard", label: "Dashboard" },
  { href: "/projects", label: "Projects" },
];

export default function Sidebar() {
  const pathname = usePathname();
  const router = useRouter();

  return (
    <aside className="flex h-screen w-56 flex-col justify-between border-r border-slate-800 bg-slate-900 p-4">
      <div>
        <p className="mb-6 text-lg font-semibold text-slate-100">Kestrel</p>
        <nav className="space-y-1">
          {links.map((link) => (
            <Link
              key={link.href}
              href={link.href}
              className={`block rounded px-3 py-2 text-sm ${
                pathname.startsWith(link.href)
                  ? "bg-slate-800 text-emerald-400"
                  : "text-slate-300 hover:bg-slate-800"
              }`}
            >
              {link.label}
            </Link>
          ))}
        </nav>
      </div>
      <button
        onClick={() => {
          clearToken();
          router.push("/login");
        }}
        className="rounded border border-slate-700 px-3 py-2 text-sm text-slate-300 hover:bg-slate-800"
      >
        Sign out
      </button>
    </aside>
  );
}
