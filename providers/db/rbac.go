package db

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
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

	Tables []TableAccess

	UserRole struct {
		ID RoleID `db:"id"`

		CreatedAt time.Time `db:"created_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedAt time.Time `db:"updated_at"`
		UpdatedBy string    `db:"updated_by"`

		Description *string `db:"description"`
		DatabaseID  ID      `db:"database_id"`
		Tables      Tables  `json:"tables"`
	}

	TableAccess struct {
		Name       string     `json:"name"`
		Columns    []string   `json:"columns"`
		AccessType AccessType `json:"access_type"`
	}
)

func (a *UserRole) CanReadTable(name string, columns ...string) bool {
	for _, table := range a.Tables {
		if table.Name == name {
			// TODO check columns here
			return true
		}
	}
	return false
}

func (t Tables) Value() (driver.Value, error) {
	j, err := json.Marshal(t)
	return j, err
}

func (t *Tables) Scan(src any) error {
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("type assertion .([]byte) failed")
	}

	var i Tables
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*t = i
	return nil
}
