CREATE SCHEMA provider_db;
SET search_path TO provider_db;

CREATE TABLE IF NOT EXISTS "user_role" (
    id TEXT PRIMARY KEY,

    created_at TIMESTAMP WITH TIME ZONE,
    created_by TEXT,
    updated_AT TIMESTAMP WITH TIME ZONE,
    updated_by TEXT,

    description TEXT,
    database_id TEXT,
    tables JSONB
);

CREATE INDEX user_role_db_id_idx ON user_role(database_id);
CREATE INDEX user_role_db_id_role_id_idx ON user_role(database_id, id);
