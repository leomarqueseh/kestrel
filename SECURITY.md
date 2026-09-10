# Security Policy

## Authorized use only

Kestrel is a tool for conducting security assessments against targets you **own** or have **explicit written authorization** to test — your own lab environment, CTF platforms, bug bounty programs in scope, or a client engagement with a signed authorization.

Kestrel does not include, and will not include, features designed to:
- Scan or attack targets without configured, explicit authorization
- Automatically weaponize a finding into a working exploit
- Perform destructive testing (data deletion, service disruption) by default

The target scope model (Phase 05) is designed to make out-of-scope operations a rejected request, not just a documented rule.

## Finding lifecycle

Every finding in Kestrel moves through a controlled state machine — scanner output is a *signal*, not a confirmed result:

```
Detected → Needs Validation → Confirmed
```

Only a `Confirmed` finding, with attached evidence (request, response, timestamp, tester notes), is included in a final report as a validated result.

## Initial threat model (Phase 00)

| Asset | Threat | Mitigation (planned) |
|---|---|---|
| Target scope data | Scope tampering leads to unauthorized testing | Scope enforced server-side, audit-logged |
| Stored credentials for targets/integrations | Leak grants attacker access to real assets | Secrets manager / encrypted storage, never in git |
| Evidence & findings data | Sensitive vulnerability data exposed | RBAC (`admin`/`analyst`/`viewer`), auth required on all endpoints |
| Reports | Leaked report exposes a client's real vulnerabilities | Access-controlled generation & storage |
| Kestrel's own codebase | Vulnerable dependency or injected code | Dependency scanning, SAST, secret scanning in CI (Phase 14) |

This table will be revisited and expanded as each phase introduces new components (Phase 19 — Security Hardening is a full pass).

## Reporting a vulnerability in Kestrel itself

This is a learning/portfolio project in early development. If you find a security issue in the Kestrel codebase itself, please open a GitHub issue describing it — avoid including working exploit code in a public issue while the project has no user base to protect yet.
