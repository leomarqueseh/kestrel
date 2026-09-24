<div align="center">

# 🦅 KESTREL

**Security Assessment & Penetration Testing Platform**

*Observe. Analyze. Validate.*

**More than a scanner — a complete Security Engineering platform.**

[![Status](https://img.shields.io/badge/status-active%20development-orange)]()
[![Go](https://img.shields.io/badge/core-Go-00ADD8?logo=go)]()
[![PostgreSQL](https://img.shields.io/badge/database-PostgreSQL-336791?logo=postgresql)]()
[![Next.js](https://img.shields.io/badge/frontend-Next.js-black?logo=next.js)]()
[![Python](https://img.shields.io/badge/automation-Python-3776AB?logo=python)]()
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Roadmap](https://img.shields.io/badge/roadmap-v2.0-blueviolet)]()

</div>

---

## ⚠️ Authorized use only

Kestrel is built to be used **exclusively** against targets you own or have **explicit written permission** to test — including authorized labs, CTF environments, bug bounty programs, and your own infrastructure.

Authorization and scope are treated as **core security controls**, not optional configuration.

Kestrel is designed around:

* Explicit target authorization
* Scope enforcement
* Allow/deny rules
* Out-of-scope protection
* Role-based access control
* Auditable assessment activity

See [`SECURITY.md`](SECURITY.md) for the full security policy and responsible disclosure process.

---

## 🎯 Why Kestrel

Most portfolio pentest projects are a single script that runs a scan and prints results.

Kestrel is a **real security assessment platform** with a working, end-to-end lifecycle: authorization → discovery → detection → validated evidence → reporting.

The platform focuses on:

* **Scope-first assessment** with enforced authorization boundaries
* **Detection ≠ Validation ≠ Exploitation**
* Persistent **attack surface inventory**
* Evidence-driven findings
* Reproducible security assessments
* Auditable assessment activity
* Modular security architecture
* Automated reporting
* A working web dashboard
* Future DevSecOps and cloud integration

The goal is not simply to find something suspicious.

The goal is to answer:

> **What was observed, why does it matter, can it be validated, what evidence supports it, and how should it be remediated?**

### Assessment lifecycle

```text
Authorization
      │
      ▼
Scope & Target
      │
      ▼
Discovery
      │
      ▼
Enumeration
      │
      ▼
Detection
      │
      ▼
Potential Finding
      │
      ▼
Security Validation
      │
 ┌────┴────┐
 ▼         ▼
Rejected  Confirmed
            │
            ▼
         Evidence
            │
            ▼
           Impact
            │
            ▼
          Report
            │
            ▼
       Remediation
            │
            ▼
          Retest
```

> **Detection ≠ Validation ≠ Exploitation**

A detection represents an observation or potential issue.
Validation determines whether the observation is actually reproducible and relevant.
Exploitation, where applicable and explicitly authorized, is a separate security activity.

**This full lifecycle is implemented and working today** — from an authenticated target being scanned, through NVD-correlated detection, manual validation with evidence, to a generated report.

---

## 📌 Status

🚧 **Active development**

Kestrel has a working backend covering the entire assessment lifecycle, plus a web dashboard currently being finished.

### Current progress

* ✅ Phase 00 — Planning
* ✅ Phase 01 — Project Foundation
* ✅ Phase 02 — Backend Architecture
* ✅ Phase 03 — Database & Migrations
* ✅ Phase 04 — Authentication & RBAC
* ✅ Phase 05 — Target Management & Scope Enforcement
* ✅ Phase 06 — Reconnaissance
* ✅ Phase 07 — Enumeration & Attack Surface Inventory
* ✅ Phase 08 — Vulnerability Assessment (NVD-correlated)
* ✅ Phase 09 — Security Validation (evidence-backed confirmation)
* ✅ Phase 10 — Reporting (HTML / JSON; PDF intentionally deferred)
* 🚧 Phase 11 — Frontend Dashboard (in progress)

See [`CHANGELOG.md`](CHANGELOG.md) for detailed development history.

---

# ✨ Features

## Implemented

| Category                     | Capability                                       | Status |
| ----------------------------- | ------------------------------------------------ | :----: |
| **Authentication**            | JWT access & refresh tokens                      |    ✅   |
| **Authorization**             | `admin` / `analyst` / `viewer` RBAC               |    ✅   |
| **Target Management**         | Authorized target lifecycle                       |    ✅   |
| **Scope Enforcement**         | Authorization-aware assessment boundaries         |    ✅   |
| **Reconnaissance**            | DNS resolution                                    |    ✅   |
| **Reconnaissance**            | Passive subdomain discovery via `crt.sh`          |    ✅   |
| **Reconnaissance**            | HTTP service detection                            |    ✅   |
| **Enumeration**               | Concurrent TCP port scanning                       |    ✅   |
| **Enumeration**               | Passive banner grabbing                            |    ✅   |
| **Attack Surface**            | Asset aggregation across scan runs                 |    ✅   |
| **Vulnerability Assessment**  | NVD-correlated CVE detection with real CVSS scores |    ✅   |
| **Finding Lifecycle**         | `detected → needs_validation → confirmed/rejected` enforced in code and DB |    ✅   |
| **Evidence**                  | Request/response/notes attached to confirmed findings |    ✅   |
| **Reporting**                 | HTML report (styled, printable) and JSON export    |    ✅   |
| **Dashboard API**              | Aggregate counts by severity and status            |    ✅   |
| **Web Dashboard**              | Login, projects, targets, scans, findings UI       |   🚧   |
| **Persistence**               | PostgreSQL-backed assessment data                  |    ✅   |

## In development / planned

| Category           | Capability                               | Status |
| ------------------- | ----------------------------------------- | :----: |
| **Reporting**        | PDF export                                |    ⏳   |
| **Automation**       | Python security automation                |    ⏳   |
| **DevSecOps**        | SAST / SCA / secret scanning / SBOM       |    ⏳   |
| **Infrastructure**   | Docker / Kubernetes / Terraform            |    ⏳   |
| **Cloud**            | AWS deployment                             |    ⏳   |
| **Observability**    | Logs / metrics / traces                    |    ⏳   |
| **Testing**          | Security regression / fuzzing / E2E        |    ⏳   |
| **Documentation**    | OpenAPI / ADRs / runbooks / threat model   |    ⏳   |

---

# 📚 Standards & References

Kestrel's security assessment methodology is designed around established industry standards and security frameworks.

| Standard                      | Purpose                           |
| ------------------------------ | ---------------------------------- |
| **OWASP Top 10**               | Web application security risks     |
| **OWASP API Security Top 10**  | API-specific security risks        |
| **OWASP WSTG**                 | Web security testing methodology   |
| **CWE**                        | Weakness classification            |
| **CVSS**                       | Vulnerability severity assessment (sourced live from the NVD) |
| **PTES**                       | Penetration testing methodology    |

These standards provide structure and vocabulary for the platform. They do not replace professional security analysis or assessment judgment.

---

## 🔎 Vulnerability taxonomy

Kestrel's vulnerability engine is designed to map findings to established security classifications.

Examples include:

| Category                       | Examples                                       |
| -------------------------------- | ----------------------------------------------- |
| **Broken Access Control**        | IDOR, BOLA, BFLA                                 |
| **Security Misconfiguration**    | Security headers, CORS, exposed files            |
| **Supply Chain Security**        | Vulnerable dependencies, CI/CD weaknesses        |
| **Cryptographic Failures**       | Weak TLS, insecure hashing, exposed secrets      |
| **Injection**                    | XSS, SQL Injection, command injection, LFI       |
| **Insecure Design**              | Business logic flaws, race conditions            |
| **Authentication Failures**      | Session handling, authentication controls        |
| **Integrity Failures**           | Unsafe deserialization, integrity validation     |
| **Logging & Monitoring**         | Insufficient logging and alerting                |
| **Exceptional Conditions**       | Improper error handling and resilience issues    |

---

# 🏗️ Architecture

Kestrel follows a **modular monolith** architecture.

Instead of introducing microservices prematurely, the platform uses a single deployable Go service with clearly separated internal domains, plus an independent Next.js frontend that consumes the API exclusively.

```text
                         ┌──────────────────────┐
                         │       Client         │
                         │ Web (Next.js) / API  │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │       Handler        │
                         │     HTTP / API       │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │       Service        │
                         │    Business Logic    │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │       Domain         │
                         │ Security Concepts     │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │     Repository       │
                         │   Data Persistence    │
                         └──────────┬───────────┘
                                    │
                                    ▼
                         ┌──────────────────────┐
                         │      PostgreSQL      │
                         └──────────────────────┘
```

### Internal security domains

```text
internal/
├── auth/        # authentication, RBAC middleware
├── project/     # assessment projects
├── target/      # scope and target authorization
├── scan/        # scan orchestration (recon, enumeration)
├── asset/       # attack surface inventory
├── recon/       # DNS, subdomain discovery, HTTP probing
├── enum/        # concurrent TCP port enumeration
├── nvd/         # NVD CVE correlation client
├── finding/     # vulnerability findings + validation workflow
├── evidence/    # proof attached to confirmed findings
├── report/      # HTML/JSON report generation
└── dashboard/   # aggregate security posture summary

web/             # Next.js frontend, consumes the API exclusively
```

The architecture intentionally separates security concerns so that new capabilities can be added without turning Kestrel into a collection of tightly coupled scanning scripts.

Full architectural decisions and diagrams are documented in [`ARCHITECTURE.md`](ARCHITECTURE.md).

---

# 🧩 Security Architecture

Authorization is part of the assessment workflow.

A target should not automatically become eligible for testing simply because it exists in the database.

```text
┌─────────────────┐
│ Target Created  │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│  Unauthorized   │
└────────┬────────┘
         │
         │ Explicit approval (admin only)
         ▼
┌─────────────────┐
│   Authorized    │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ Assessment      │
│ Allowed         │
└─────────────────┘
```

The security model is enforced across multiple boundaries — every layer below is implemented and active, not aspirational:

```text
Authentication      → JWT, verified on every protected route
       │
       ▼
Authorization       → RBAC (admin/analyst/viewer)
       │
       ▼
Target Scope        → unauthorized targets never reach the network
       │
       ▼
Assessment Permission → only admin can approve a target for testing
       │
       ▼
Security Operation  → recon, enumeration, and assessment modules
       │
       ▼
Audit / Evidence    → confirmed findings require attached evidence
```

---

# 🗺️ Roadmap — v2.0

<details open>
<summary><strong>Block 1 · Foundation</strong></summary>

| Phase | Focus                                                    | Status |
| ----- | ---------------------------------------------------------| :----: |
| 00    | Planning — objectives, scope, threat model, architecture |    ✅   |
| 01    | Project foundation — repository, boilerplate, `/health`  |    ✅   |

</details>

<details open>
<summary><strong>Block 2 · Platform Core</strong></summary>

| Phase | Focus                                                            | Status |
| ----- | ----------------------------------------------------------------- | :----: |
| 02    | Backend architecture — handler → service → domain → repository    |    ✅   |
| 03    | Database & migrations — PostgreSQL, entities, relationships       |    ✅   |
| 04    | Authentication — JWT, refresh tokens, password hashing, RBAC      |    ✅   |
| 05    | Target management — authorization and scope enforcement           |    ✅   |

</details>

<details open>
<summary><strong>Block 3 · Security Engine</strong></summary>

| Phase | Focus                                                          | Status |
| ----- | ----------------------------------------------------------------| :----: |
| 06    | Reconnaissance — DNS, subdomains, HTTP discovery                |    ✅   |
| 07    | Enumeration — ports, services, attack surface inventory          |    ✅   |
| 08    | Vulnerability assessment — NVD correlation, real CVSS scoring    |    ✅   |
| 09    | Security validation — potential finding → confirmed/rejected     |    ✅   |
| 10    | Reporting — methodology, findings, evidence, remediation (HTML/JSON) |    ✅   |

</details>

<details open>
<summary><strong>Block 4 · Interface & Automation</strong></summary>

| Phase | Focus                      | Status |
| ----- | --------------------------- | :----: |
| 11    | Frontend dashboard          |   🚧   |
| 12    | Python security automation  |    ⏳   |

</details>

<details>
<summary><strong>Block 5 · Delivery & Scale</strong></summary>

| Phase | Focus                                                            | Status |
| ----- | ------------------------------------------------------------------| :----: |
| 13    | Docker — production images, non-root, health checks               |    ⏳   |
| 14    | DevSecOps — SAST, SCA, secrets, SBOM, provenance                  |    ⏳   |
| 15    | Kubernetes — deployment, networking, policies, scaling            |    ⏳   |
| 16    | AWS — VPC, EKS, RDS, S3, IAM, CloudWatch                          |    ⏳   |
| 17    | Terraform — infrastructure as code                                |    ⏳   |
| 18    | Observability — logs, metrics, traces                             |    ⏳   |
| 19    | Security hardening — application, API, DB, containers             |    ⏳   |
| 20    | Testing — unit, integration, E2E, security regression, fuzzing    |    ⏳   |
| 21    | Documentation — API, ADRs, threat model, runbooks                 |    ⏳   |
| 22    | Production-like deployment — CI/CD, monitoring, rollback          |    ⏳   |
| 23    | Portfolio — releases, demo, screenshots, case studies             |    ⏳   |

</details>

---

# 🧰 Tech Stack

| Layer                   | Technology                   | Purpose                                           |
| ------------------------ | ----------------------------- | --------------------------------------------------|
| **Core API**              | Go                             | Core platform, concurrency and security workflows |
| **HTTP Router**           | chi                            | Lightweight HTTP routing and middleware            |
| **Database**              | PostgreSQL                     | Persistent relational data                          |
| **Database Driver**       | pgx                             | PostgreSQL access from Go                            |
| **Authentication**        | JWT                             | Access and refresh token authentication              |
| **Password Hashing**      | bcrypt                          | Password protection                                  |
| **Migrations**            | golang-migrate                  | Version-controlled schema migrations                 |
| **Vulnerability Data**    | NVD API                         | Live CVE correlation with CVSS scoring               |
| **Report Rendering**      | Go `html/template`              | XSS-safe HTML report generation                       |
| **Frontend**              | React + Next.js + TypeScript + Tailwind | Web dashboard, consumes the API exclusively |
| **Security Automation**   | Python (planned)                | Future automation and specialized security tooling    |
| **Containers**            | Docker / Docker Compose         | Local PostgreSQL today; full app containerization planned |
| **CI/CD**                 | GitHub Actions (planned)        | Future security and delivery pipeline                 |
| **Cache / Queue**         | Redis (planned)                 | Introduced when orchestration requires it              |
| **Orchestration**         | Kubernetes (planned)            | Future container orchestration                         |
| **Infrastructure**        | Terraform (planned)             | Future infrastructure as code                           |
| **Cloud**                 | AWS (planned)                   | Future cloud deployment                                 |
| **Observability**         | Prometheus / Grafana (planned)  | Future monitoring stack                                 |

> Kestrel intentionally avoids introducing infrastructure complexity before the application requires it. Redis, Kubernetes, AWS, and Terraform remain deferred to later roadmap phases.

---

# 🔬 Attack Surface Inventory

One of Kestrel's core concepts is maintaining an aggregated view of the assets discovered during security assessments — implemented and queryable today via `GET /targets/{id}/assets`, joining results across every scan ever run against a target.

```text
Target
 │
 ├── Domains
 │    ├── example.com
 │    ├── api.example.com
 │    └── dev.example.com
 │
 ├── IP Addresses
 │
 ├── Services
 │    ├── HTTP
 │    ├── HTTPS
 │    └── SSH
 │
 ├── Ports
 │    ├── 22
 │    ├── 80
 │    └── 443
 │
 └── Technologies
      ├── Web Server
      ├── Framework
      └── Application
```

Instead of treating every scan as an isolated execution, Kestrel builds a persistent view of the target's observable attack surface over time.

---

# 🚀 Getting Started

## Requirements

* Go 1.21+
* Node.js 18+ (for the web dashboard)
* Docker
* Docker Compose
* [golang-migrate](https://github.com/golang-migrate/migrate)

### Installing prerequisites

Already have these installed? Skip to [Backend](#backend) below.

<details>
<summary><strong>Debian / Ubuntu / Kali Linux</strong></summary>

**Go 1.21+** — the version in `apt` is often outdated, so install from the official archive:

```bash
cd /tmp
wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.23.4.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version
```

> Check [go.dev/dl](https://go.dev/dl/) for the current release if `go1.23.4` is no longer the latest.

**Node.js 18+:**

```bash
node --version   # many distros already ship a recent enough version
```

If missing or older than 18:

```bash
curl -fsSL https://deb.nodesource.com/setup_lts.x | sudo -E bash -
sudo apt install -y nodejs
node --version
```

**Docker + Docker Compose:**

```bash
sudo apt update
sudo apt install docker.io -y
sudo systemctl enable --now docker
sudo usermod -aG docker $USER
newgrp docker
docker --version
docker compose version
```

**golang-migrate** (requires Go to be installed first):

```bash
export PATH=$PATH:$(go env GOPATH)/bin
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate -version
```

</details>

<details>
<summary><strong>macOS (Homebrew)</strong></summary>

```bash
brew install go node golang-migrate
```

Docker on macOS requires [Docker Desktop](https://www.docker.com/products/docker-desktop/) rather than the CLI-only package — install it separately and make sure it's running before `docker compose up`.

</details>

<details>
<summary><strong>Windows</strong></summary>

Use [WSL2](https://learn.microsoft.com/windows/wsl/install) with an Ubuntu distribution and follow the Debian/Ubuntu instructions above inside it. Install [Docker Desktop](https://www.docker.com/products/docker-desktop/) with WSL2 integration enabled rather than installing Docker directly inside WSL.

</details>

## Backend

```bash
git clone https://github.com/leomarqueseh/kestrel.git
cd kestrel

cp .env.example .env
docker compose up -d

export DATABASE_URL="postgres://kestrel:kestrel@localhost:5432/kestrel?sslmode=disable"
migrate -database "$DATABASE_URL" -path migrations up

go mod tidy
go run ./cmd/seed -email admin@kestrel.local -password "change-me"
make run
```

API available at `http://localhost:8080`.

```bash
curl http://localhost:8080/health
```

Expected response:

```json
{"status":"ok","service":"kestrel-api","time":"...","database":"ok"}
```

## Frontend

```bash
cd web
cp .env.local.example .env.local   # NEXT_PUBLIC_API_URL=http://localhost:8080
npm install
npm run dev
```

Dashboard available at `http://localhost:3000`.

---

# 📁 Project Structure

```text
kestrel/
├── cmd/
│   ├── api/              # API entrypoint
│   └── seed/             # bootstraps the first admin user
├── docs/                 # architecture and data-model documentation
├── internal/
│   ├── auth/              # authentication and RBAC
│   ├── project/           # project management
│   ├── target/            # target and scope management
│   ├── scan/               # scan orchestration
│   ├── asset/              # attack surface assets
│   ├── recon/               # reconnaissance
│   ├── enum/                # port enumeration
│   ├── nvd/                  # NVD CVE correlation client
│   ├── finding/               # vulnerability findings + validation
│   ├── evidence/               # evidence attached to findings
│   ├── report/                  # report generation
│   ├── dashboard/                # aggregate summary
│   └── platform/postgres/         # DB connection pool
├── migrations/            # database migrations
├── web/                    # Next.js frontend
├── automation/              # planned Python automation
├── Makefile
├── ARCHITECTURE.md
├── CHANGELOG.md
├── CONTRIBUTING.md
├── LICENSE
├── README.md
└── SECURITY.md
```

---

# 🔐 Security Testing Philosophy

Kestrel is designed around a simple principle:

> **A finding is more than a scanner output.**

A mature assessment establishes:

```text
Observation
    ↓
Context
    ↓
Validation
    ↓
Evidence
    ↓
Impact
    ↓
Remediation
    ↓
Retest
```

This principle is enforced in code, not just in documentation: a finding cannot become `confirmed` in Kestrel without evidence attached in the same request.

---

# 🤝 Contributing

Kestrel is currently following a structured development roadmap.

Contributions, technical discussions, bug reports, and architectural feedback are welcome.

Before contributing:

1. Read [`CONTRIBUTING.md`](CONTRIBUTING.md).
2. Understand the current roadmap.
3. Keep changes scoped to the relevant module.
4. Add or update tests where appropriate.
5. Document security-sensitive behavior.
6. Never introduce functionality that bypasses authorization or scope controls.

---

# 📄 License

Kestrel is released under the [MIT License](LICENSE).

---

# 🔒 Security

Found a vulnerability in Kestrel itself?

Please review [`SECURITY.md`](SECURITY.md) for the responsible disclosure process.

Do not publicly disclose sensitive vulnerability details before the issue has been responsibly reported and evaluated.

---

<div align="center">

### 🦅 Kestrel

**Observe. Analyze. Validate.**

*Construído para um futuro mais seguro.*

</div>
