# Architecture

## Guiding principle

**Simplicity → Correctness → Security → Scalability.**

No technology is added because it looks impressive on a resume. Each one is added when there is a concrete technical reason for it, and that reason is documented here.

## Phase 1: Modular Monolith (current target)

```
                    ┌─────────────────┐
                    │    Frontend     │
                    │ React / Next.js │
                    └────────┬────────┘
                             │ HTTPS / REST
                             ▼
                    ┌─────────────────┐
                    │    REST API     │
                    │       Go        │
                    └────────┬────────┘
                             │
              ┌──────────────┼──────────────┐
              ▼              ▼              ▼
        ┌──────────┐   ┌──────────┐   ┌──────────┐
        │ Scanner  │   │   Auth   │   │ Reports  │
        │  Engine  │   │ Service  │   │ Service  │
        └────┬─────┘   └──────────┘   └──────────┘
             │
             ▼
        ┌──────────┐
        │PostgreSQL│
        └──────────┘
```

**Why a monolith first:** target management, auth, scanning orchestration, and reporting are tightly coupled in the early phases — they share the same data model and change together. Splitting them into services now would mean distributed-systems complexity (network calls, retries, eventual consistency) with no corresponding benefit. Each internal module (`internal/scanner`, `internal/auth`, `internal/reports`) has a clear boundary and its own package, so splitting later — if there's ever a real reason to — is a refactor, not a rewrite.

**Go = core platform.** Handles the HTTP API, domain logic, persistence, auth, and orchestrates when scans run.

**Python = security automation module**, invoked by the Scanner Engine for tasks where its ecosystem is strongest (parsing tool output, recon scripts, HTTP-level testing). It never duplicates responsibilities owned by Go (auth, persistence, business rules).

## Request flow

```
HTTP request
   ↓
Handler          (parses/validates the request)
   ↓
Service          (business logic, orchestration)
   ↓
Domain           (entities, rules)
   ↓
Repository       (data access, interface-based for testability)
   ↓
PostgreSQL
```

## Phase 2: Cloud-native (later, once justified)

```
                       AWS
                        │
                 ┌──────┴──────┐
                 │ Load Balancer│
                 └──────┬──────┘
                        │
                    Kubernetes
                        │
        ┌───────────────┼───────────────┐
        ▼               ▼               ▼
      API            Worker          Scanner
        │               │               │
        └───────────────┼───────────────┘
                        │
                 ┌──────┴──────┐
                 │ PostgreSQL  │
                 │    Redis    │
                 │    Queue    │
                 └─────────────┘
```

This stage is introduced in later roadmap phases (Docker → Kubernetes → AWS → Terraform → Observability), once there's a real need to run scans concurrently at scale, queue long-running jobs, and deploy to a production-like environment. Redis and a queue are added when scan orchestration actually requires them — not preemptively.

## Core data model (initial entities)

```
users
projects
targets      — scoped, authorized assets only
scans
assets       — discovered during enumeration
findings     — detected → needs validation → confirmed
evidence     — attached to confirmed findings
reports
```

## Key architectural decisions

| Decision | Rationale |
|---|---|
| Modular monolith, not microservices | No current justification for network-boundary complexity |
| `findings.status` state machine (`detected → needs_validation → confirmed`) | Scanner output is never treated as a confirmed vulnerability |
| Target scope enforced at the API layer | Operations outside authorized scope must be rejected, not just discouraged |
| Python isolated to `automation/` | Prevents logic duplication between Go and Python |
| PostgreSQL only, no Redis yet | Avoids adding infrastructure before there's a workload that needs it |
