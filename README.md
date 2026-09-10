# Kestrel

**Kestrel** is a security assessment platform that guides an authorized penetration test end-to-end: reconnaissance, enumeration, vulnerability assessment, validated evidence collection, and automated reporting.

> ⚠️ **Authorized use only.** Kestrel is built to be used against targets you own or have explicit written permission to test (labs, CTFs, bug bounty programs, your own infrastructure). See [SECURITY.md](./SECURITY.md) for the full scope policy.

## Why

Most portfolio pentest projects are a single script that runs a scan and prints results. Kestrel is built as a real, modular platform: it manages scope, separates *detected* findings from *confirmed* findings, keeps an audit trail of evidence, and produces a report a client could actually receive.

## Status

🚧 Early development — Phase 00 (Planning) complete, Phase 01 (Project Foundation) in progress.
See [CHANGELOG.md](./CHANGELOG.md) for progress and the roadmap below for what's next.

## Features (planned)

- Scoped target management (domains, IPs, URLs, authorized CIDR ranges)
- Reconnaissance module (DNS, subdomains, HTTP services)
- Enumeration module (ports, services, versions, technologies) → attack surface inventory
- Vulnerability assessment with CVSS-based severity and a `detected → needs validation → confirmed` workflow
- Evidence collection (request/response, timestamps, tester notes) for every confirmed finding
- Automated report generation (executive summary, methodology, findings, evidence, recommendations) as HTML/PDF/JSON
- Role-based access control (`admin`, `analyst`, `viewer`)

## Tech stack

| Layer | Choice | Why |
|---|---|---|
| Core API | Go (Chi/Gin) | Core platform: performance, strong typing, concurrency for scan orchestration |
| Security automation | Python | Recon/parsing scripts where its ecosystem is unmatched — never duplicates Go's responsibilities |
| Database | PostgreSQL | Relational integrity for targets, scans, findings, evidence |
| Cache/Queue | Redis | Added when scan orchestration actually needs it (not from day one) |
| Frontend | React + Next.js + TypeScript | Consumes the API exclusively |
| Containers | Docker / Docker Compose | Local + CI parity |
| CI/CD | GitHub Actions | Lint → tests → SAST → dependency scan → secret scan → build |

Kubernetes, Terraform, and AWS are introduced later in the roadmap, once there is a concrete reason to orchestrate and deploy at that scale — see [ARCHITECTURE.md](./ARCHITECTURE.md).

## Architecture

Kestrel starts as a **modular monolith** — one deployable Go service with clearly separated internal modules (Scanner Engine, Auth Service, Reports Service) — rather than microservices from day one. Full rationale and diagrams: [ARCHITECTURE.md](./ARCHITECTURE.md).

## Roadmap

| Phase | Focus |
|---|---|
| 00 | Planning ✅ |
| 01 | Project foundation (repo structure, health endpoint) |
| 02 | Backend architecture (handler → service → domain → repository) |
| 03 | Database & migrations |
| 04 | Authentication & RBAC |
| 05 | Target management (scope enforcement) |
| 06–07 | Reconnaissance & enumeration |
| 08–09 | Vulnerability assessment & validation |
| 10 | Reporting |
| 11 | Frontend |
| 12 | Python security automation |
| 13–23 | Docker → DevSecOps → Kubernetes → AWS → Terraform → Observability → Hardening → Testing → Docs → Deployment → Portfolio polish |

## Getting started

_Coming in Phase 01 — the API doesn't exist yet._

## License

[MIT](./LICENSE)

## Security

Found a security issue in Kestrel itself? See [SECURITY.md](./SECURITY.md) for how to report it responsibly.
