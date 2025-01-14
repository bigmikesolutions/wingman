package db

import "errors"

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
		ID             RoleID           `json:"id"`
		Description    *string          `db:"description"`
		DatabaseAccess []DatabaseAccess `db:"database_access"`
	}

	RBAC interface {
		Add(role UserRole) error
		Check(id ID, role RoleID) (*DatabaseAccess, error)
	}

	InMemoryRBAC struct {
		roles map[RoleID]*UserRole
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

func (i InMemoryRBAC) Add(role UserRole) error {
	i.roles[role.ID] = &role
	return nil
}

func (i InMemoryRBAC) Check(id ID, roleID RoleID) (*DatabaseAccess, error) {
	role, ok := i.roles[roleID]
	if !ok {
		return nil, nil
	}

	for _, access := range role.DatabaseAccess {
		if access.DatabaseID == id {
			clone := access
			return &clone, nil
		}
	}

	return nil, nil
}
