package iam

type (
	AuthType int
	UserID     string
	AccessType int
)

const (
	AuthTypeGcp AuthType = iota
	AuthTypeAws
	AuthTypeHashicorpVault

	AccessTypeRead = iota
	AccessTypeCreate
	AccessTypeDelete
	AccessTypeUpdate
)

type User interface {
	GetID() UserID
	GetName() string
	GetAuthType() AuthType
	GetPolicies() []Policy
}

type Policy interface {
	GetResource() Resource
	GetAccess() []AccessType
}

type Resource interface {
	GetRoot() Resource
	GetName() string
}

