package actions

import (
	"context"
	"fmt"

	"github.com/bigmikesolutions/wingman/pkg/iam/access"
	"github.com/bigmikesolutions/wingman/pkg/iam/identity"

	"github.com/bigmikesolutions/wingman/pkg/cqrs"
	"github.com/bigmikesolutions/wingman/pkg/provider"
)

func AuthorizeUserResourceAccess(
	ctx context.Context,
	queryBus cqrs.QueryBus,
	providerID provider.ID,
	resourceType provider.ResourceType,
	resourceID provider.ResourceID,
	requiredAccess access.AccessType,
	additionalRequiredAccesses ...access.AccessType,
) error {
	userSession := identity.GetUserSession(ctx)
	if userSession == nil {
		resourceIDValue := ""
		if resourceID != "" {
			resourceIDValue = string(resourceID)
		}
		return identity.NewUserSessionNotFoundError(
			fmt.Sprintf("provider id: %s, resource type: %s, resource id: %s",
				providerID, resourceType, resourceIDValue),
		)
	}
	result, err := queryBus.ExecuteQuery(ctx, &access.QueryGetUserResourceAccessRights{
		UserID:       userSession.User().Id(),
		ProviderID:   providerID,
		ResourceType: resourceType,
		ResourceID:   resourceID,
	})
	if err != nil {
		return err
	}
	userAccessRights := result.(access.QueryGetUserAccessPolicyResult)
	if noAccessErr := access.ContainsAccessRight(userAccessRights.AccessTypes, requiredAccess); noAccessErr != nil {
		return noAccessErr
	}
	for _, addReqAccess := range additionalRequiredAccesses {
		if noAccessErr := access.ContainsAccessRight(userAccessRights.AccessTypes, addReqAccess); noAccessErr != nil {
			return noAccessErr
		}
	}
	return nil
}
