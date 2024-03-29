package mock

import (
	"context"

	"github.com/bigmikesolutions/wingman/core/cqrs"
	"github.com/bigmikesolutions/wingman/core/iam/access"
)

type QueryGetUserAccessPolicyHandler struct{}

func (h *QueryGetUserAccessPolicyHandler) GetType() cqrs.QueryType {
	return access.QueryTypeGetUserResourceAccessRights
}

func (h *QueryGetUserAccessPolicyHandler) Handle(ctx context.Context, q cqrs.Query) (interface{}, error) {
	query := q.(*access.QueryGetUserResourceAccessRights)
	return access.QueryGetUserAccessPolicyResult{
		ProviderID:   query.ProviderID,
		UserID:       query.UserID,
		ResourceID:   query.ResourceID,
		ResourceType: query.ResourceType,
		AccessTypes: []access.AccessType{
			access.AccessTypeRead,
			access.AccessTypeCreate,
			access.AccessTypeUpdate,
			access.AccessTypeDelete,
		},
	}, nil
}
