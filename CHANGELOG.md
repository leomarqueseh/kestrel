# Changelog

All notable changes to this project are documented here.
Format loosely follows [Keep a Changelog](https://keepachangelog.com/).

## [Unreleased]

### Phase 07 — Enumeration
- Added concurrent TCP port scanner (`internal/enum`) over common ports with passive banner grabbing
- Refactored `scan.Service` into a shared lifecycle (`run()`) reused by recon and enumeration
- Added Attack Surface Inventory: `GET /targets/{id}/assets` aggregates results across every scan

### Phase 06 — Reconnaissance
- Added `internal/recon`: DNS resolution, passive subdomain discovery via crt.sh, HTTP service probing
- Added `scan` and `asset` domains with a pending → running → completed/failed lifecycle
- Enforced scope at the service layer: unauthorized targets are rejected before any network request

### Phase 05 — Target Management
- Added minimal project management (create/list, owned by the authenticated user)
- Added target CRUD scoped to a project
- Targets are created `authorized: false` by default; only `admin` can authorize them

### Phase 04 — Authentication
- Added login, JWT access/refresh tokens, bcrypt password hashing
- Added RBAC middleware (`RequireAuth`, `RequireRole`) for `admin`/`analyst`/`viewer`
- Added `cmd/seed` to bootstrap the first admin user
- Rate-limited the login endpoint against brute force

### Phase 03 — Database
- Added PostgreSQL schema via golang-migrate: users, projects, targets, scans, assets, findings, evidence, reports
- Connected the API to PostgreSQL via a pgx connection pool
- Extended `/health` to report database connectivity

### Phase 02 — Backend Architecture
- Introduced layered architecture (handler → service) via `internal/health` and `internal/version`
- Wired routing through chi with logging and panic-recovery middleware

### Phase 01 — Project Foundation
- Set up Go module and standard project layout (`cmd/`, `internal/`, `pkg/`, `automation/`)
- Added `GET /health` endpoint

### Phase 00 — Planning
- Defined project name, objective, scope, and tech stack
- Wrote initial architecture (modular monolith, Go core + Python automation module)
- Wrote initial threat model and authorized-use policy
- Created README, ARCHITECTURE, SECURITY, CONTRIBUTING, LICENSE, CHANGELOG
