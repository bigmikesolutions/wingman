package iam

type (
	ResourceID string
)

type Resource interface {
	GetID() ResourceID
	GetName() string
	GetType() string
	Execute(action ResourceAction) error
}
