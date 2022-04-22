package iam

type (
	GroupID string
)

type Group struct {
	id       GroupID
	name     string
	policies []ResourceAccessPolicy
	users    []User `gorm:"many2many:user_groups;"`
}

func NewGroup(id GroupID, name string, policies []ResourceAccessPolicy, users []User) *Group {
	return &Group{id: id, name: name, policies: policies, users: users}
}

func (g Group) Id() GroupID {
	return g.id
}

func (g Group) Name() string {
	return g.name
}

func (g Group) Policies() []ResourceAccessPolicy {
	return g.policies
}

func (g Group) Users() []User {
	return g.users
}
