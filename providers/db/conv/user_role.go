package conv

import (
	"github.com/bigmikesolutions/wingman/graphql/model"
	"github.com/bigmikesolutions/wingman/providers/db/rbac"
)

func InternalAccessType(v *model.AccessType) *rbac.AccessType {
	if v == nil {
		return nil
	}

	var r rbac.AccessType
	switch *v {
	case model.AccessTypeReadOnly:
		r = rbac.ReadOnlyAccess

	case model.AccessTypeWriteOnly:
		r = rbac.WriteOnlyAccess

	case model.AccessTypeReadWrite:
		r = rbac.ReadWriteAccess

	default:
		return nil
	}

	return &r
}
