package api

import (
	"context"
	"fmt"

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

func (s *HTTPServer) EnvironmentQuery(ctx context.Context, env string) (*model.Environment, error) {
	query := `
query($env: EnvironmentID!) {
	environment(id: $env) {
		id
		description
		createdAt
		modifiedAt
	}
}
`

	vars := map[string]any{
		"env": env,
	}

	var response environmentResponse
	err := s.graphqlExecute(ctx, query, vars, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Errors) > 0 {
		return nil, fmt.Errorf("graphql error: %+v", response.Errors)
	}

	return response.Data.Environment, nil
}
