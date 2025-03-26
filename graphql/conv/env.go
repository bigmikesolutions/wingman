package conv

import (
	"github.com/bigmikesolutions/wingman/graphql/model"
	"github.com/bigmikesolutions/wingman/server/env"
)

func EnvInternalToPublic(v *env.Environment) *model.Environment {
	if v == nil {
		return nil
	}

	return &model.Environment{
		ID:          v.ID,
		Description: v.Description,
	}
}

func CreateEnvironmentInputToInternal(v model.CreateEnvironmentInput) env.Environment {
	return env.Environment{
		ID:          v.Env,
		Description: v.Description,
	}
}
