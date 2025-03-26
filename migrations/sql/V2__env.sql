CREATE SCHEMA wingman;
SET search_path TO wingman;

CREATE TABLE IF NOT EXISTS "environments" (
    id TEXT PRIMARY KEY,

    created_at TIMESTAMP WITH TIME ZONE,
    created_by TEXT,
    updated_at TIMESTAMP WITH TIME ZONE,
    updated_by TEXT,

    description TEXT
);
