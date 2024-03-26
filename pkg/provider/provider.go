package provider

import "github.com/bigmikesolutions/wingman/pkg/cqrs"

type Factory = func() (cqrs.Config, error)
