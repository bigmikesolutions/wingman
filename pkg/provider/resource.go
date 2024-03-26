package provider

type (
	ResourceID   string
	ResourceType string
	ID           string
)

type Resource interface {
	GetID() ResourceID
	GetType() ResourceType
	GetProviderID() ID
	GetName() string
	Execute(action ResourceAction) error
	GetInfo() interface{}
}
