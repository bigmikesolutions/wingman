package env

import "time"

type (
	ID = string

	Environment struct {
		ID    ID     `db:"id"`
		OrgID string `db:"org_id"`

		CreatedAt time.Time `db:"created_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedAt time.Time `db:"updated_at"`
		UpdatedBy string    `db:"updated_by"`

		Description *string `db:"description"`
	}
)
