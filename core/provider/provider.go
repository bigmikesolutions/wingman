package provider

import "github.com/bigmikesolutions/wingman/core/cqrs"

type Factory = func() (cqrs.Config, error)

type Provider interface {
	ID() ID
	HandleQuery(query QueryGetResource)
}
