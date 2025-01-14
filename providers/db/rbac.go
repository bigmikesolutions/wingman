package db

type (
	UserID = string

	UserRole struct {
		ID          string `json:"id"`
		Read        bool   `db:"read"`
		Description string `db:"description"`
		DatabaseIDs []ID   `db:"database_ids"`
	}

	RBAC interface {
		CanRead(id ID, user UserID) bool
	}
)
