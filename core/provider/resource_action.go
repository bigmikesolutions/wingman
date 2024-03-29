package provider

import (
	"context"
)

type ResourceAction interface {
	// GetResourceAccessSession() iam.ResourceAccessSession
	GetActionParams() map[string]interface{}
	GetContext() context.Context
}
