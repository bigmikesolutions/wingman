package iam

type (
	UserID string
)

type User struct {
	id       UserID
	name     string
	policies []ResourceAccessPolicy
	groups   []Group `gorm:"many2many:user_groups;"`
}

func NewUser(id UserID, name string, policies []ResourceAccessPolicy, groups []Group) *User {
	return &User{id: id, name: name, policies: policies, groups: groups}
}

func (u User) Id() UserID {
	return u.id
}

func (u User) Name() string {
	return u.name
}

func (u User) Policies() []ResourceAccessPolicy {
	return u.policies
}

func (u User) Groups() []Group {
	return u.groups
}
