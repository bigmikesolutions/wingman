package db

import (
	"context"
	"errors"
	"time"
)

type AccessType uint8

var (
	ErrDatabaseAccessDenied = errors.New("database access denied")
	ErrTableAccessDenied    = errors.New("table access denied")
)

const (
	ReadOnlyAccess AccessType = iota
	WriteOnlyAccess
	ReadWriteAccess
)

type (
	RoleID = string

	RBAC interface {
		CreateUserRole(ctx context.Context, role UserRole) error
		FindUserRolesByDatabaseID(ctx context.Context, id ID) ([]UserRole, error)
		Close() error
	}

	DatabaseAccess struct {
		DatabaseID ID
		Tables     []TableAccess
	}

	TableAccess struct {
		Name       string
		Columns    []string
		AccessType AccessType
	}

	UserRole struct {
		ID RoleID `db:"id"`

		CreatedAt time.Time `db:"created_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedAt time.Time `db:"updated_at"`
		UpdatedBy string    `db:"updated_by"`

		Description    *string          `db:"description"`
		DatabaseID     ID               `db:"database_id"`
		DatabaseAccess []DatabaseAccess `db:"tables"`
	}
)

func (db *DatabaseAccess) CanReadTable(name string, columns ...string) bool {
	for _, table := range db.Tables {
		if table.Name == name {
			// TODO check columns here
			return true
		}
	}
	return false
}
