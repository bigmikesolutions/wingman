package provider

import "github.com/bigmikesolutions/wingman/core/cqrs"

type Factory = func() (cqrs.Config, error)
