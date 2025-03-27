package repo

const (
	createEnv = `
INSERT INTO wingman.environments(
	id, org_id,
	created_at, created_by, updated_at, updated_by,
	description
)
VALUES(
    :id, :org_id,
    :created_at, :created_by, :updated_at, :updated_by,
	:description
)`

	selectEnvByID = `
SELECT
*
FROM
	wingman.environments
WHERE 
  	org_id = $1
	and id = $2
`
)
