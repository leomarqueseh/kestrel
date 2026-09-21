-- Enables gen_random_uuid() for UUID primary keys
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- People who can access the platform
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'analyst' CHECK (role IN ('admin', 'analyst', 'viewer')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- A named assessment/engagement that groups targets and scans
CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    description TEXT,
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Authorized assets in scope for a project (domain, IP, URL or lab CIDR)
CREATE TABLE targets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    value TEXT NOT NULL,
    target_type TEXT NOT NULL CHECK (target_type IN ('domain', 'ip', 'url', 'cidr')),
    authorized BOOLEAN NOT NULL DEFAULT false,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- A single execution of a module (recon, enumeration...) against a target
CREATE TABLE scans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    target_id UUID NOT NULL REFERENCES targets(id) ON DELETE CASCADE,
    module TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'running', 'completed', 'failed')),
    started_at TIMESTAMPTZ,
    finished_at TIMESTAMPTZ,
    error TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Discovered ports/services/technologies from enumeration
CREATE TABLE assets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    scan_id UUID NOT NULL REFERENCES scans(id) ON DELETE CASCADE,
    host TEXT NOT NULL,
    port INTEGER,
    protocol TEXT,
    service TEXT,
    version TEXT,
    technology TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Potential or confirmed vulnerabilities tied to an asset
CREATE TABLE findings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id UUID NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT,
    severity TEXT NOT NULL CHECK (severity IN ('critical', 'high', 'medium', 'low', 'informational')),
    cvss NUMERIC(3,1),
    cwe TEXT,
    status TEXT NOT NULL DEFAULT 'detected' CHECK (status IN ('detected', 'needs_validation', 'confirmed', 'false_positive')),
    recommendation TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Proof attached to a confirmed finding
CREATE TABLE evidence (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    finding_id UUID NOT NULL REFERENCES findings(id) ON DELETE CASCADE,
    request TEXT,
    response TEXT,
    notes TEXT,
    captured_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Generated report output for a project
CREATE TABLE reports (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    format TEXT NOT NULL CHECK (format IN ('html', 'pdf', 'json')),
    generated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Indexes for the lookups the application will do most often
CREATE INDEX idx_targets_project_id ON targets(project_id);
CREATE INDEX idx_scans_target_id ON scans(target_id);
CREATE INDEX idx_assets_scan_id ON assets(scan_id);
CREATE INDEX idx_findings_asset_id ON findings(asset_id);
CREATE INDEX idx_findings_status ON findings(status);
CREATE INDEX idx_evidence_finding_id ON evidence(finding_id);
CREATE INDEX idx_reports_project_id ON reports(project_id);
