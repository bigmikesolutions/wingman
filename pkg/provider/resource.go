package provider

type (
	ResourceID   string
	ResourceType string
	ProviderID   string
)

type Resource interface {
	GetID() ResourceID
	GetType() ResourceType
	GetProviderID() ProviderID
	GetName() string
	Execute(action ResourceAction) error
	GetInfo() interface{}
}
