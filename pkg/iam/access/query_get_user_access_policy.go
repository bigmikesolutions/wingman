package access

import (
	"github.com/bigmikesolutions/wingman/pkg/cqrs"
	"github.com/bigmikesolutions/wingman/pkg/iam/identity"
	"github.com/bigmikesolutions/wingman/pkg/provider"
)

const (
	QueryTypeGetUserResourceAccessRights cqrs.QueryType = "AccessGetUserResourceAccessRights"
)

type QueryGetUserResourceAccessRights struct {
	ProviderID   provider.ID
	ResourceType provider.ResourceType
	ResourceID   provider.ResourceID
	UserID       identity.UserID
}

func (c QueryGetUserResourceAccessRights) GetType() cqrs.QueryType {
	return QueryTypeGetUserResourceAccessRights
}

type QueryGetUserAccessPolicyResult struct {
	ProviderID   provider.ID
	ResourceType provider.ResourceType
	ResourceID   provider.ResourceID
	UserID       identity.UserID
	AccessTypes  []AccessType
}
