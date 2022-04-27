package provider

import "github.com/bigmikesolutions/wingman/pkg/cqrs"

type ProviderFactory = func(*cqrs.Config) error
