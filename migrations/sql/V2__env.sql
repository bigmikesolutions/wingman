CREATE SCHEMA wingman;
SET search_path TO wingman;

CREATE TABLE IF NOT EXISTS "environments" (
    id      TEXT NOT NULL,
    org_id  TEXT NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE,
    created_by TEXT,
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by TEXT,

    description TEXT,

    PRIMARY KEY (id, org_id)
);
