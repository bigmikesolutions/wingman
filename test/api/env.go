package api

import (
	"context"

	"github.com/bigmikesolutions/wingman/graphql/model"
)

func (s *HTTPServer) EnvGrantMutation(
	ctx context.Context,
	input model.EnvGrantInput,
) (*model.EnvGrantPayload, error) {
	var mutation struct {
		model.EnvGrantPayload `graphql:"envGrant(input: $input)"`
	}

	err := s.graphql.Mutate(ctx, &mutation, map[string]interface{}{
		"input": input,
	})
	if err != nil {
		return nil, err
	}

	return &mutation.EnvGrantPayload, nil
}

func (s *HTTPServer) CreateEnvironment(
	ctx context.Context,
	input model.CreateEnvironmentInput,
) (*model.CreateEnvironmentPayload, error) {
	var mutation struct {
		model.CreateEnvironmentPayload `graphql:"createEnvironment(input: $input)"`
	}

	err := s.graphql.Mutate(ctx, &mutation, map[string]interface{}{
		"input": input,
	})
	if err != nil {
		return nil, err
	}

	return &mutation.CreateEnvironmentPayload, nil
}
