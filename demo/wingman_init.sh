#!/bin/bash
set -e

PGPASSWORD=pass psql -v ON_ERROR_STOP=1 -h localhost -p 5432 --username "admin" --dbname "wingman" <<-EOSQL

DROP DATABASE "test-db-1";
CREATE DATABASE "test-db-1";
GRANT ALL PRIVILEGES ON DATABASE "test-db-1" TO admin;

SET search_path TO wingman;

DELETE FROM environments;

INSERT INTO
  environments(id, org_id, description, created_at, created_by, updated_by)
VALUES
  ('test', 'bms', 'demo', now(), 'scripts', '');

SET search_path TO provider_db;

DELETE FROM user_role;

INSERT INTO
  user_role(id, org_id, env, database_id, info, tables, created_at, created_by, updated_by)
VALUES
  ('test', 'bms', 'test', 'test-db-1', 'read_only', '[{"Name":"students", "Columns":[], "AccessType":"read_only"}]', now(), 'scripts', '');


EOSQL

PGPASSWORD=pass psql -v ON_ERROR_STOP=1 -h localhost -p 5432 --username "admin" --dbname "test-db-1" <<-EOSQL

CREATE TABLE IF NOT EXISTS students (
	id SERIAL,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	age INT NOT NULL
);

DELETE FROM students;

INSERT INTO students(id,first_name,last_name,age) values
	('1','johny','bravo',30),
	('2','mike','tyson',51),
	('3','pamela','anderson',65);

EOSQL

export VAULT_ADDR=http://127.0.0.1:8200
vault login root
vault kv put \
  secret/providers/db/organisations/bms/environments/test/connections/test-db-1 \
    id="test-db-1" \
    env="test-env" \
    org_id="bms" \
    driver="pgx" \
    host="localhost" \
    name="test-db-1" \
    port=5432 \
    user="admin" \
    pass="pass"
