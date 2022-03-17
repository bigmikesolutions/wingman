package iam

type (
	UserID         string
	GroupID        string
	AccessPolicyID string
	AccessType     int
)

const (
	AccessTypeRead = iota
	AccessTypeCreate
	AccessTypeDelete
	AccessTypeUpdate
)

type User interface {
	GetID() UserID
	GetName() string
	GetPolicies() []AccessPolicy
	GetGroups() []Group
}

type Group interface {
	GetID() GroupID
	GetName() string
	GetPolicies() []AccessPolicy
	GetUsers() []User
}

type AccessPolicy interface {
	GetID() AccessPolicyID
	GetName() string
	GetResource() []Resource
	GetAccess() []AccessType
}
