package iam

import "context"

type ResourceAction interface {
	GetResourceAccessSession() ResourceAccessSession
	GetActionParams() map[string]interface{}
	GetContext() context.Context
}
