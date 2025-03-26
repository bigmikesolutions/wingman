package repo

const (
	createEnv = `
INSERT INTO wingman.environments(
	id, 
	created_at, created_by, updated_at, updated_by,
	description
)
VALUES(
    :id,
    :created_at, :created_by, :updated_at, :updated_by,
	:description
)`

	selectEnvByID = `
SELECT
*
FROM
	wingman.environments
WHERE 
	id = $1
`
)
