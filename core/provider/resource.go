package provider

type (
	ResourceID   string
	ResourceType string
	ID           string
)

type Resource interface {
	GetProviderID() ID
	GetID() ResourceID
	GetType() ResourceType
	GetName() string
	Execute(action ResourceAction) error
	GetInfo() interface{}
}
