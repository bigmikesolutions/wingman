package iam

import "time"

type (
	OrganisationID string
	RoleID         string
	GroupID        string
	UserID         string

	ChangeMetaInfo struct {
		CreatedAt time.Time `db:"created_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedAt time.Time `db:"updated_at"`
		UpdatedBy string    `db:"updated_by"`
	}

	UserRole struct {
		OrgID       OrganisationID `db:"org_id"`
		ID          RoleID         `db:"id"`
		Description string         `db:"description"`

		ChangeMetaInfo
	}

	UserGroup struct {
		OrgID OrganisationID `db:"org_id"`
		ID    GroupID        `db:"id"`

		Description string   `db:"description"`
		Roles       []RoleID `db:"roles"`

		ChangeMetaInfo
	}

	User struct {
		OrgID OrganisationID `db:"org_id"`
		ID    UserID         `db:"id"`

		Email     string `db:"email"`
		FirstName string `db:"first_name"`
		LastName  string `db:"last_name"`
		Active    bool   `db:"active"`

		Roles []UserRoleBinding

		ChangeMetaInfo
	}

	UserRoleBinding struct {
		OrgID  OrganisationID `db:"org_id"`
		UserID UserID         `db:"user_id"`

		Roles  []RoleID  `db:"roles"`
		Groups []GroupID `db:"groups"`

		ChangeMetaInfo
	}
)
