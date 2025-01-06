package api

import (
	"context"
	"fmt"

	"github.com/bigmikesolutions/wingman/graphql/model/cursor"

	"github.com/bigmikesolutions/wingman/graphql/model"
)

func (s *HTTPServer) Database(ctx context.Context, env string, id string) (*model.Database, error) {
	query := `
query($env: EnvironmentID!, $databaseID:String!) {
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

func (s *HTTPServer) DatabaseTableData(
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
			driver
			table(name: $tableName, first: $first, after: $after, where: $where) {
				connectionInfo {
					endCursor
					hasNextPage
				}
				edges {
					cursor
					node {
						ts
						rows {
							index 
							values
						}
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

	return response.Data.Environment.Database, nil
}
