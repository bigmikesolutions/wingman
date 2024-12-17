package api

import (
	"context"
	"fmt"

	"github.com/bigmikesolutions/wingman/graphql/model"
)

func (s *HTTPServer) Database(ctx context.Context, env string, id string) (*model.Database, error) {
	query := `
query($env: EnvironmentID!, $databaseID: String!) {
	environment(id: $env) {
		database(id: $databaseID) {
			id
			driver
		}
	}
}
`

	vars := map[string]any{
		"env":        env,
		"databaseID": id,
	}

	var response environmentResponse
	err := s.graphqlExecute(ctx, query, vars, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Errors) > 0 {
		return nil, fmt.Errorf("graphql error: %+v", response.Errors)
	}

	return response.Data.Environment.Database, nil
}
