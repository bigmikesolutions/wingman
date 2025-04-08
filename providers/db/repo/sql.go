package repo

const (
	createUserRole = `
INSERT INTO provider_db.user_role(
	id,
	org_id, env, database_id,
	created_at, created_by, updated_at, updated_by,
	description, info, tables
)
VALUES(
    :id,
    :org_id, :env, :database_id,
    :created_at, :created_by, :updated_at, :updated_by,
	:description, :info, :tables
)`

	selectUserRolesByDatabaseByFullID = `
SELECT
*
FROM
	provider_db.user_role
WHERE 
	org_id = $1 and env = $2 and database_id = $3
`
)
