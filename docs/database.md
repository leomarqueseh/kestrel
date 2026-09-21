# Database model

PostgreSQL, 8 core entities, managed via `golang-migrate` (see `migrations/`).

## Entities

| Table | Purpose |
|---|---|
| `users` | People who access the platform (role: admin/analyst/viewer) |
| `projects` | A named assessment, owned by a user |
| `targets` | Authorized assets in scope for a project (domain/IP/URL/CIDR) |
| `scans` | An execution of a module (recon, enumeration...) against a target |
| `assets` | Ports/services/technologies discovered by a scan |
| `findings` | Vulnerabilities tied to an asset — lifecycle: detected → needs_validation → confirmed → false_positive |
| `evidence` | Proof (request/response/notes) attached to a finding |
| `reports` | Generated report output for a project |

## Relationships

## Key decisions

- UUID primary keys (`gen_random_uuid()`) instead of serial integers — avoids leaking sequential IDs (how many targets/findings exist) through the API.
- `findings.status` is constrained by a `CHECK`, enforcing the detected → confirmed lifecycle at the database level, not just in application code.
- All deletions cascade downward from `targets` (deleting a target removes its scans, assets, findings, evidence) — but `projects.owner_id` uses `ON DELETE RESTRICT`, so a user can't be deleted while still owning projects.
