package api

import (
	"context"
	"fmt"

	"github.com/bigmikesolutions/wingman/graphql/model"
	"github.com/bigmikesolutions/wingman/graphql/model/cursor"
)

func (s *HTTPServer) AddDatabaseMutation(
	ctx context.Context,
	input model.AddDatabaseInput,
) (*model.AddDatabasePayload, error) {
	var mutation struct {
		model.AddDatabasePayload `graphql:"addDatabase(input: $input)"`
	}

	err := s.graphql.Mutate(ctx, &mutation, map[string]interface{}{
		"input": input,
	})
	if err != nil {
		return nil, err
	}

	return &mutation.AddDatabasePayload, nil
}

func (s *HTTPServer) DatabaseInfoQuery(ctx context.Context, env string, id string) (*model.Database, error) {
	query := `
query($env: EnvironmentID!, $databaseID:String!) {
	environment(id: $env) {
		database(id: $databaseID) {
			id
			info {
				id
                host
                port
                driver
            }
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

	if response.Data.Environment == nil {
		return nil, nil
	}

	return response.Data.Environment.Database, nil
}

func (s *HTTPServer) DatabaseTableDataQuery(
	ctx context.Context,
	env string,
	id string,
	tableName string,
	first int,
	after cursor.Cursor,
	where model.TableFilter,
) (*model.Database, error) {
	query := `
query($env: EnvironmentID!, $databaseID: String!, $tableName: String!, $first: Int=50, $after: Cursor, $where: TableFilter) {
	environment(id: $env) {
		database(id: $databaseID) {
			id
			table(name: $tableName, first: $first, after: $after, where: $where) {
				connectionInfo {
					endCursor
					hasNextPage
				}
				edges {
					cursor
					node {
						index 
						values
					}
				}
			}
		}
	}
}
`

	vars := map[string]any{
		"env":        env,
		"databaseID": id,
		"tableName":  tableName,
		"first":      first,
		"after":      after,
		"where":      where,
	}

	var response environmentResponse
	err := s.graphqlExecute(ctx, query, vars, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Errors) > 0 {
		return nil, fmt.Errorf("graphql error: %+v", response.Errors)
	}

	if response.Data.Environment == nil {
		return nil, nil
	}

	return response.Data.Environment.Database, nil
}

func (s *HTTPServer) CreateDatabaseUserRoleMutation(
	ctx context.Context,
	input model.AddDatabaseUserRoleInput,
) (*model.AddDatabaseUserRolePayload, error) {
	var mutation struct {
		model.AddDatabaseUserRolePayload `graphql:"addDatabaseUserRole(input: $input)"`
	}

	err := s.graphql.Mutate(ctx, &mutation, map[string]interface{}{
		"input": input,
	})
	if err != nil {
		return nil, err
	}

	return &mutation.AddDatabaseUserRolePayload, nil
}
