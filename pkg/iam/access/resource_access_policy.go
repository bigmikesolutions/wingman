package access

import "github.com/bigmikesolutions/wingman/pkg/provider"

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
	resource    provider.Resource
	accessTypes []AccessType
}

func NewAccessPolicy(id AccessPolicyID, name string, resource provider.Resource, accessTypes ...AccessType) *ResourceAccessPolicy {
	return &ResourceAccessPolicy{id: id, name: name, resource: resource, accessTypes: accessTypes}
}

func (a ResourceAccessPolicy) Id() AccessPolicyID {
	return a.id
}

func (a ResourceAccessPolicy) Name() string {
	return a.name
}

func (a ResourceAccessPolicy) Resource() provider.Resource {
	return a.resource
}

func (a ResourceAccessPolicy) AccessTypes() []AccessType {
	return a.accessTypes
}

func ContainsAccessRight(accessTypes []AccessType, requiredAccess AccessType) error {
	for _, access := range accessTypes {
		if access == requiredAccess {
			return nil
		}
	}
	return NewUnauthorizedError("Access denied")
}
