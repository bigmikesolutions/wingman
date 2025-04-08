package rbac

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type (
	AccessType string
	RoleID     = string

	TableScope struct {
		Name       string     `json:"name"`
		Columns    []string   `json:"columns"`
		AccessType AccessType `json:"access_type"`
	}

	TablesScopes []TableScope

	UserRole struct {
		ID    RoleID `db:"id"`
		OrgID string `db:"org_id"`
		Env   string `db:"env"`

		CreatedAt time.Time `db:"created_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedAt time.Time `db:"updated_at"`
		UpdatedBy string    `db:"updated_by"`

		Description *string `db:"description"`
		DatabaseID  string  `db:"database_id"`

		Info   *AccessType  `json:"info"`
		Tables TablesScopes `json:"tables"`
	}
)

const (
	ReadOnlyAccess  AccessType = "read_only"
	WriteOnlyAccess            = "write_only"
	ReadWriteAccess            = "read_write"
)

func (a *UserRole) CanReadTable(name string, columns ...string) bool {
	// TODO add checks for columns here
	for _, table := range a.Tables {
		if table.Name == name {
			// TODO check columns here
			return true
		}
	}
	return false
}

func (t TablesScopes) Value() (driver.Value, error) {
	j, err := json.Marshal(t)
	return j, err
}

func (t *TablesScopes) Scan(src any) error {
	source, ok := src.([]byte)
	if !ok {
		return fmt.Errorf("type assertion .([]byte) failed")
	}

	var i TablesScopes
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*t = i
	return nil
}
