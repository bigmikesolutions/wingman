CREATE TABLE IF NOT EXISTS "db_user_role" (
    id TEXT PRIMARY KEY,
    role_id TEXT PRIMARY KEY,

    created_at TIMESTAMP WITH TIME ZONE,
    created_by TEXT,
    updated_AT TIMESTAMP WITH TIME ZONE,
    updated_by TEXT,

    description TEXT,
    database_id TEXT,
    tables JSOB,
);

CREATE INDEX db_user_role_db_id_idx ON db_user_role(database_id);
CREATE INDEX db_user_role_role_id_idx ON db_user_role(role_id);
CREATE INDEX db_user_role_db_id_role_id_idx ON db_user_role(database_id, role_id);
