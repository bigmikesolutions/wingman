package iam

type (
	AccessPolicyID string
	AccessType     int
)

const (
	AccessTypeRead = iota
	AccessTypeCreate
	AccessTypeDelete
	AccessTypeUpdate
)

type ResourceAccessPolicy struct {
	id          AccessPolicyID
	name        string
	resource    Resource
	accessTypes []AccessType
}

func NewAccessPolicy(id AccessPolicyID, name string, resource Resource, accessTypes ...AccessType) *ResourceAccessPolicy {
	return &ResourceAccessPolicy{id: id, name: name, resource: resource, accessTypes: accessTypes}
}

func (a ResourceAccessPolicy) Id() AccessPolicyID {
	return a.id
}

func (a ResourceAccessPolicy) Name() string {
	return a.name
}

func (a ResourceAccessPolicy) Resource() Resource {
	return a.resource
}

func (a ResourceAccessPolicy) AccessTypes() []AccessType {
	return a.accessTypes
}
