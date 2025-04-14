package env

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	ID             = string
	OrganisationID = string

	Environment struct {
		ID    ID             `db:"id"`
		OrgID OrganisationID `db:"org_id"`

		CreatedAt time.Time  `db:"created_at"`
		CreatedBy string     `db:"created_by"`
		UpdatedAt *time.Time `db:"updated_at"`
		UpdatedBy *string    `db:"updated_by"`

		Description *string `db:"description"`
	}
)

func (v Environment) Validate() error {
	return validation.ValidateStruct(&v,
		validation.Field(&v.ID, validation.Required, validation.Length(3, 255)),
		validation.Field(&v.OrgID, validation.Required, validation.Length(3, 255)),
		validation.Field(&v.Description, validation.NilOrNotEmpty),
	)
}
