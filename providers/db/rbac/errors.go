package rbac

import "errors"

var (
	ErrDatabaseInfoDenied   = errors.New("database info denied")
	ErrDatabaseAccessDenied = errors.New("database access denied")
	ErrTableAccessDenied    = errors.New("table access denied")
)
