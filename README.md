<div align="center">

# 🦅 KESTREL

**Security Assessment & Penetration Testing Platform**

*Observe. Analyze. Validate.*

More than a scanner — a complete Security Engineering platform.

[![Status](https://img.shields.io/badge/status-early%20development-orange)]()
[![Go](https://img.shields.io/badge/core-Go-00ADD8?logo=go)]()
[![Python](https://img.shields.io/badge/automation-Python-3776AB?logo=python)]()
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Roadmap](https://img.shields.io/badge/roadmap-v2.0-blueviolet)]()

</div>

---

## ⚠️ Authorized use only

Kestrel is built to be used **exclusively** against targets you own or have **explicit written permission** to test — labs, CTFs, bug bounty programs, or your own infrastructure. Scope validation, allow/deny lists, and out-of-scope protection are core platform features, not an afterthought. See [`SECURITY.md`](SECURITY.md) for the full scope policy and responsible disclosure process.

---

## 🎯 Why Kestrel

Most portfolio pentest projects are a single script that runs a scan and prints results. Kestrel is built as a real, modular platform:

- Manages **scope** with enforced authorization boundaries
- Separates **detected** findings from **validated/confirmed** findings
- Keeps an **audit trail** of evidence (requests, responses, timestamps, tester notes)
- Produces a **report a client could actually receive** — not a terminal dump

> "Técnica + Automação + Inteligência = Segurança Real"

## 📌 Status

🚧 **Early development** — Phase 00 (Planning) complete, Phase 01 (Project Foundation) in progress.
See [`CHANGELOG.md`](CHANGELOG.md) for detailed progress.

---

## ✨ Features (planned)

| Category | Capability |
|---|---|
| **Target Management** | Scoped domains, IPs, URLs, and authorized CIDR ranges with allow/deny lists |
| **Reconnaissance** | DNS, subdomain enumeration, HTTP service and certificate discovery |
| **Enumeration** | Ports, services, versions, technologies → full attack surface inventory |
| **Vulnerability Assessment** | OWASP Top 10 / API Top 10 / WSTG-aligned detection with CVSS scoring |
| **Validation Workflow** | `Potential Finding → Validation → Confirmed / False Positive` with confidence tracking |
| **Evidence Collection** | Request/response captures, timestamps, tester notes for every confirmed finding |
| **Reporting** | Automated executive summary, methodology, findings, evidence, remediation — HTML / PDF / JSON |
| **Access Control** | RBAC (admin, analyst, viewer), audit logs, rate limiting |
| **Dashboard** | Web frontend for projects, targets, scans, assets, findings, and reports |

---

## 📚 Standards & References

Kestrel's assessment methodology is built directly on top of industry standards, not ad-hoc heuristics:

- **OWASP Top 10:2025** — Web application risk categories
- **OWASP API Security Top 10** — API-specific risk categories
- **OWASP WSTG** (Web Security Testing Guide) — Testing methodology
- **CWE** (Common Weakness Enumeration) — Weakness classification
- **CVSS** (Common Vulnerability Scoring System) — Severity scoring

<details>
<summary><strong>Vulnerability catalog (OWASP Top 10:2025 mapping)</strong></summary>

| ID | Category | Examples |
|---|---|---|
| A01 | Broken Access Control | IDOR, BOLA, BFLA |
| A02 | Security Misconfiguration | Headers, CORS, `.git` exposure |
| A03 | Supply Chain Failures | Dependencies, CI/CD |
| A04 | Cryptographic Failures | TLS, hashing, secrets |
| A05 | Injection | XSS, SQLi, SSRF, LFI |
| A06 | Insecure Design | Business logic, race conditions |
| A07 | Authentication Failures | Brute force, session handling |
| A08 | Data Integrity Failures | Deserialization, integrity checks |
| A09 | Logging & Alerting Failures | Logs, monitoring |
| A10 | Exceptional Conditions | Error handling, DoS |

</details>

---

## 🏗️ Architecture

Kestrel starts as a **modular monolith** — one deployable Go service with clearly separated internal modules (Scanner Engine, Auth Service, Reports Service) — rather than microservices from day one. Full rationale and diagrams: [`ARCHITECTURE.md`](ARCHITECTURE.md).

The build is organized into five blocks:

| Block | Focus |
|---|---|
| **1 · Foundation** | Project scope, threat model, architecture, repo boilerplate |
| **2 · Platform Core** | Backend architecture, database, authentication, target management |
| **3 · Security Engine** | Reconnaissance → enumeration → vulnerability assessment → validation → reporting |
| **4 · Interface & Automation** | Frontend dashboard, Python security automation |
| **5 · Delivery / Scale** | Docker, DevSecOps, Kubernetes, AWS, Terraform, observability, hardening, testing, docs |

---

## 🧰 Tech stack

| Layer | Choice | Why |
|---|---|---|
| **Core API** | Go (Chi/Gin) | Performance, strong typing, concurrency for scan orchestration |
| **Security automation** | Python | Recon/parsing scripts where its ecosystem is unmatched — never duplicates Go's responsibilities |
| **Database** | PostgreSQL | Relational integrity for targets, scans, findings, evidence |
| **Cache/Queue** | Redis | Added when scan orchestration actually needs it (not from day one) |
| **Frontend** | React + Next.js + TypeScript | Consumes the API exclusively |
| **Containers** | Docker / Docker Compose | Local + CI parity |
| **Orchestration** | Kubernetes | Introduced once there's a concrete reason to orchestrate at scale |
| **IaC** | Terraform | Infra as code — VPC, EKS, IAM, RDS, S3, load balancer |
| **Cloud** | AWS | VPC, EKS, RDS, ElastiCache, S3, IAM, CloudWatch |
| **CI/CD** | GitHub Actions | Lint → tests → SAST → dependency scan → secret scan → build → deploy |

Kubernetes, Terraform, and AWS are introduced later in the roadmap, once there is a concrete reason to orchestrate and deploy at that scale — see [`ARCHITECTURE.md`](ARCHITECTURE.md).

---

## 🗺️ Roadmap (v2.0)

<details open>
<summary><strong>Block 1 · Foundation</strong></summary>

| Phase | Focus | Status |
|---|---|---|
| 00 | Planning — objective, scope, threat model, architecture | ✅ |
| 01 | Project foundation — repo structure, boilerplate, `/health` endpoint | 🚧 |

</details>

<details>
<summary><strong>Block 2 · Platform Core</strong></summary>

| Phase | Focus |
|---|---|
| 02 | Backend architecture (handler → service → domain → repository, SOLID, DI) |
| 03 | Database & migrations (PostgreSQL, entities, relationships, indices) |
| 04 | Authentication (JWT, refresh tokens, password hashing, RBAC, audit logs) |
| 05 | Target management (scope validation, allow/deny lists, CIDR/domain/IP/URL) |

</details>

<details>
<summary><strong>Block 3 · Security Engine</strong></summary>

| Phase | Focus |
|---|---|
| 06 | Reconnaissance (DNS, subdomains, HTTP, tech discovery, certificates) |
| 07 | Enumeration (ports, services, versions, attack surface inventory) |
| 08 | Vulnerability assessment (OWASP Top 10:2025 / API Top 10 / WSTG, CWE, CVSS) |
| 09 | Security validation (potential finding → validation → confirmed/rejected) |
| 10 | Reporting (executive summary, scope, methodology, findings, evidence) |

</details>

<details>
<summary><strong>Block 4 · Interface & Automation</strong></summary>

| Phase | Focus |
|---|---|
| 11 | Frontend dashboard (projects, targets, scans, assets, findings, reports) |
| 12 | Python security automation (recon, parsers, HTTP, reporting utils) |

</details>

<details>
<summary><strong>Block 5 · Delivery / Scale</strong></summary>

| Phase | Focus |
|---|---|
| 13 | Docker (multi-stage, non-root, healthcheck, secrets) |
| 14 | DevSecOps (SAST, SCA, secret scan, SBOM, provenance) |
| 15 | Kubernetes (Deploy, Service, Ingress, ConfigMap/Secret, HPA, RBAC, network policy) |
| 16 | AWS (VPC, EKS, RDS, ElastiCache, S3, IAM, CloudWatch, load balancer) |
| 17 | Terraform (IaC, VPC, EKS, RDS, S3, load balancer) |
| 18 | Observability (logs, metrics, traces, Prometheus, Grafana, dashboards) |
| 19 | Security hardening (app/API/DB, Docker/K8s, threat modeling, TLS) |
| 20 | Testing (unit, integration, API, E2E, security regression, fuzzing) |
| 21 | Documentation (architecture, API/OpenAPI, threat model, ADRs, runbooks) |
| 22 | Production-like deployment (CI/CD, monitor, rollback, health checks, auto-scaling) |
| 23 | Portfolio (GitHub, demo, releases, screenshots, case studies) |

</details>

---

## 🚀 Getting started

> Coming in Phase 01 — the API doesn't exist yet.

```bash
# Placeholder — updated once Phase 01 ships
git clone https://github.com/kestrel/kestrel.git
cd kestrel
make setup
```

## 📁 Project structure

```
kestrel/
├── cmd/                # Entrypoints
├── internal/
│   ├── handler/        # HTTP layer
│   ├── service/         # Business logic
│   ├── domain/          # Core entities
│   └── repository/      # Data access
├── automation/          # Python security automation
├── web/                 # React + Next.js frontend
├── deployments/         # Docker, Kubernetes, Terraform
└── docs/                 # Architecture, ADRs, runbooks
```

## 🤝 Contributing

This project currently follows a solo development roadmap (see [Roadmap](#️-roadmap-v20)). Issues and discussion are welcome — see [`CONTRIBUTING.md`](CONTRIBUTING.md) once published.

## 📄 License

MIT — see [`LICENSE`](LICENSE).

## 🔒 Security

Found a vulnerability in Kestrel itself, or have questions about authorized scope? See [`SECURITY.md`](SECURITY.md) for the responsible disclosure process.

---

<div align="center">
<sub>Kestrel — Construído para um futuro mais seguro.</sub>
</div>
