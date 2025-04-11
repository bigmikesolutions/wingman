package env

import (
	"context"
	"fmt"

	"github.com/bigmikesolutions/wingman/client/graphqlclient"
	"github.com/bigmikesolutions/wingman/graphql/model"
)

type Client struct {
	client *graphqlclient.Client
}

func New(client *graphqlclient.Client) *Client {
	return &Client{client: client}
}

func (s *Client) EnvGrantMutation(
	ctx context.Context,
	input model.EnvGrantInput,
) (*model.EnvGrantPayload, error) {
	var mutation struct {
		model.EnvGrantPayload `graphql:"envGrant(input: $input)"`
	}

	err := s.client.Mutate(ctx, &mutation, map[string]interface{}{
		"input": input,
	})
	if err != nil {
		return nil, err
	}

	return &mutation.EnvGrantPayload, nil
}

func (s *Client) CreateEnvironment(
	ctx context.Context,
	input model.CreateEnvironmentInput,
) (*model.CreateEnvironmentPayload, error) {
	var mutation struct {
		model.CreateEnvironmentPayload `graphql:"createEnvironment(input: $input)"`
	}

	err := s.client.Mutate(ctx, &mutation, map[string]interface{}{
		"input": input,
	})
	if err != nil {
		return nil, err
	}

	return &mutation.CreateEnvironmentPayload, nil
}

func (s *Client) EnvironmentQuery(ctx context.Context, env string) (*model.Environment, error) {
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

	var response graphqlclient.EnvironmentResponse
	err := s.client.Execute(ctx, query, vars, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Errors) > 0 {
		return nil, fmt.Errorf("graphql error: %+v", response.Errors)
	}

	return response.Data.Environment, nil
}
