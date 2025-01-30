package rbac

import "errors"

var (
	ErrDatabaseAccessDenied = errors.New("database access denied")
	ErrTableAccessDenied    = errors.New("table access denied")
)
