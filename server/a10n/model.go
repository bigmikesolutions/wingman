package a10n

const (
	// AgentRead is a read-only role for global admins
	AgentRead = "agent-read-only"
	// AgentWrite is a write-only role for global admins
	AgentWrite = "agent-write-only"

	// organisation level roles

	AdminRole          = "admin"           // can change: env, users, roles, role bindings
	ProjectManagerRole = "project_manager" // can: approve env access
	DeveloperRole      = "developer"       // can: access env if access approved by PM
)

type (
	OrganisationID string
	UserID         string
	RoleID         string

	Organisation struct {
		ID   OrganisationID `json:"id"`
		Name string         `json:"name"`
	}

	UserRoles struct {
		OrganisationID *OrganisationID
		Roles          []string
	}

	UserRoleBinding struct {
		OrganisationID *OrganisationID
		UserID         UserID
		RoleID         string
	}

	User struct {
		OrganisationID *OrganisationID
		ID             UserID
		Email          string
		Roles          []UserRoles
	}
)
