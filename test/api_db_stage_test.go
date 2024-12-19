package test

import (
	"testing"

	"github.com/bigmikesolutions/wingman/graphql/model/cursor"
	"github.com/bigmikesolutions/wingman/providers/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bigmikesolutions/wingman/graphql/model"
	"github.com/bigmikesolutions/wingman/test/api"
)

const (
	driverPGX = "pgx"
)

type ApiDatabaseStage struct {
	t      *testing.T
	server *api.HTTPServer

	queryDatabase *model.Database
	err           error
}

func NewApiDatabaseStage(t *testing.T) *ApiDatabaseStage {
	server, err := api.New()
	require.Nil(t, err, "api server")

	return &ApiDatabaseStage{
		t:      t,
		server: server,
	}
}

func (s *ApiDatabaseStage) Close() {
	s.server.Close()
}

func (s *ApiDatabaseStage) Given() *ApiDatabaseStage {
	return s
}

func (s *ApiDatabaseStage) When() *ApiDatabaseStage {
	return s
}

func (s *ApiDatabaseStage) Then() *ApiDatabaseStage {
	return s
}

func (s *ApiDatabaseStage) And() *ApiDatabaseStage {
	return s
}

func (s *ApiDatabaseStage) ServerIsUpAndRunning() *ApiDatabaseStage {
	api.AssertHeartbeat(s.t, s.server)
	return s
}

func (s *ApiDatabaseStage) QueryDatabase(
	env string,
	id string,
) *ApiDatabaseStage {
	ctx, cancel := testContext()
	defer cancel()
	s.queryDatabase, s.err = s.server.Database(ctx, env, id)
	return s
}

func (s *ApiDatabaseStage) QueryDatabaseTableData(
	env string,
	id string,
	name string,
	first int,
	after cursor.Cursor,
	where model.TableFilter,
) *ApiDatabaseStage {
	ctx, cancel := testContext()
	defer cancel()
	s.queryDatabase, s.err = s.server.DatabaseTableData(ctx, env, id, name, first, after, where)
	return s
}

func (s *ApiDatabaseStage) NoClientError() *ApiDatabaseStage {
	assert.NoError(s.t, s.err, "query has failed")
	return s
}

func (s *ApiDatabaseStage) DatabaseIsProvided(database db.ConnectionInfo) *ApiDatabaseStage {
	require.Nilf(s.t, s.server.Resolver.DB.Register(database), "register database: %+v", database)
	return s
}
