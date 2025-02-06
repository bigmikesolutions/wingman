package repo

const (
	createUserRole = `
INSERT INTO provider_db.user_role(
	id, 
	created_at, created_by, updated_at, updated_by,
	description, database_id, info, tables
)
VALUES(
    :id,
    :created_at, :created_by, :updated_at, :updated_by,
	:description, :database_id, :info, :tables
)`

	selectUserRolesByDatabaseID = `
SELECT
*
FROM
	provider_db.user_role
WHERE 
	database_id = $1
`
)
