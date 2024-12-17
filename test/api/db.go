package api

import (
	"context"

	"github.com/bigmikesolutions/wingman/graphql/model"
)

func (s *HTTPServer) Database(ctx context.Context, env string, id string) (*model.Database, error) {
	query := `
query(env: String!, databaseID: String!) {
	environment(id: $env) {
		database(id: $databaseID) {
			id
			driver
		}
	}
}
`

	var response struct {
		Environment *model.Environment `graphql:"database(id: $id)"`
	}

	vars := map[string]any{
		"env":        env,
		"databaseID": id,
	}

	err := s.graphqlExecute(ctx, query, vars, &response)
	if err != nil {
		return nil, err
	}

	return response.Environment.Database, nil
}
