package test

import (
	"fmt"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib" // drivers
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // drivers
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bigmikesolutions/wingman/graphql/model"
	"github.com/bigmikesolutions/wingman/graphql/model/cursor"
	"github.com/bigmikesolutions/wingman/providers/db"
	"github.com/bigmikesolutions/wingman/test/api"
)

const (
	sqlTableStudents         = "students"
	sqlTableNonExistingTable = "non-existing-table"

	sqlCreateTableStudents = `
CREATE TABLE students (
	id SERIAL,
	first_name TEXT NOT NULL,
	last_name TEXT NOT NULL,
	age INT NOT NULL
);
`

	sqlInsertStudents = `
INSERT INTO students(id,first_name,last_name,age) values
	('1','johny','bravo',30),
	('2','mike','tyson',51),
	('3','pamela','anderson',65);
`

	expDriverPostgres = "POSTGRES"
)

var dbCleanTables = []string{
	"provider_db.user_role",
}

type ApiDatabaseStage struct {
	t      *testing.T
	server *api.HTTPServer
	dbx    *sqlx.DB

	queryDatabase *model.Database
	err           error
}

func NewApiDatabaseStage(t *testing.T) *ApiDatabaseStage {
	dbx := mustDB()
	server, err := api.New(dbx, newProviders(dbx))
	require.Nil(t, err, "api server")

	return &ApiDatabaseStage{
		t:      t,
		server: server,
		dbx:    dbx,
	}
}

func (s *ApiDatabaseStage) Close() {
	s.server.Close()

	cleanTables(s.dbx, dbCleanTables)
	_ = s.dbx.Close()
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

func (s *ApiDatabaseStage) DatabaseInfoQuery(
	env string,
	id string,
) *ApiDatabaseStage {
	ctx, cancel := testContext()
	defer cancel()
	s.queryDatabase, s.err = s.server.DatabaseInfoQuery(ctx, env, id)
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
	s.queryDatabase, s.err = s.server.DatabaseTableDataQuery(ctx, env, id, name, first, after, where)
	return s
}

func (s *ApiDatabaseStage) NoClientError() *ApiDatabaseStage {
	assert.NoError(s.t, s.err, "query has failed")
	return s
}

func (s *ApiDatabaseStage) ClientErrorIs(expMsg string) *ApiDatabaseStage {
	require.Error(s.t, s.err, "client error expected", expMsg)
	assert.Contains(s.t, s.err.Error(), expMsg, "unexpected client error")
	return s
}

func (s *ApiDatabaseStage) DatabaseIsProvided(database db.ConnectionInfo) *ApiDatabaseStage {
	ctx, cancel := testContext()
	defer cancel()
	err := s.server.Resolver.Providers.DB.Register(ctx, database)
	require.Nilf(s.t, err, "register database: %+v", database)
	return s
}

func (s *ApiDatabaseStage) DatabaseStatement(dbID, statement string, args ...any) *ApiDatabaseStage {
	ctx, cancel := testContext()
	defer cancel()

	info, err := s.server.Resolver.Providers.DB.Info(ctx, dbID)
	require.NoError(s.t, err, "db info error")
	require.NotNil(s.t, info, "db info missing")

	conn, err := sqlx.Connect(info.Driver, fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		info.User, info.Pass,
		info.Host, info.Port,
		info.Name,
	))
	require.Nilf(s.t, err, "db connection")

	rows, err := conn.QueryxContext(ctx, statement, args...)
	require.Nilf(s.t, err, "statement error")

	rows.Next()

	return s
}

func (s *ApiDatabaseStage) DatabaseInfoIsReturned(dbID, driver string) *ApiDatabaseStage {
	require.NotNil(s.t, s.queryDatabase, "query database is nil")

	assert.Equal(s.t, dbID, s.queryDatabase.ID, "database ID")
	assert.Equal(s.t, dbID, s.queryDatabase.Info.ID, "database info: ID")
	assert.Equal(s.t, driver, string(s.queryDatabase.Info.Driver), "database info: driver")
	assert.Equal(s.t, "localhost", s.queryDatabase.Info.Host, "database info: host")
	assert.NotEqual(s.t, 0, s.queryDatabase.Info.Port, "database info: port")
	return s
}

func (s *ApiDatabaseStage) TableQueryHasNextPage(nextPage bool) *ApiDatabaseStage {
	require.NotNil(s.t, s.queryDatabase.Table, "table data must be returned")
	require.NotNil(s.t, s.queryDatabase.Table.ConnectionInfo, "connection info must be returned")

	assert.Equal(s.t, nextPage, s.queryDatabase.Table.ConnectionInfo.HasNextPage, "next page")
	if nextPage {
		assert.NotEqual(s.t, "", string(s.queryDatabase.Table.ConnectionInfo.EndCursor), "end cursor missing")
	}
	return s
}

func (s *ApiDatabaseStage) TableHasRows(expRows ...model.TableRow) *ApiDatabaseStage {
	require.NotNil(s.t, s.queryDatabase.Table, "table data must be returned")

	for _, expRow := range expRows {
		found := false
		for _, edge := range s.queryDatabase.Table.Edges {
			if *expRow.Index == *edge.Node.Index {
				// TODO enable cursor checks
				// assert.NotEqual(s.t, "", string(edge.Cursor), "cursor missing")
				assert.Equalf(s.t, len(expRow.Values), len(edge.Node.Values), "unexpeted values at row: %d", *expRow.Index)
				for idx, expValue := range expRow.Values {
					assert.Equalf(s.t, *expValue, *edge.Node.Values[idx], "row value mismatch at %d", idx)
				}
				found = true
				break
			}
		}
		assert.Truef(s.t, found, "table row not found: %d", *expRow.Index)
	}

	return s
}

func (s *ApiDatabaseStage) NoTableRows() *ApiDatabaseStage {
	require.NotNil(s.t, s.queryDatabase.Table, "table data must be returned")

	for _, edge := range s.queryDatabase.Table.Edges {
		assert.Empty(s.t, edge.Node.Values, "no table rows expected")
	}

	return s
}

func (s *ApiDatabaseStage) DatabaseUserRoleIsCreated(input model.AddDatabaseUserRoleInput) *ApiDatabaseStage {
	ctx, cancel := testContext()
	defer cancel()
	payload, err := s.server.CreateDatabaseUserRoleMutation(ctx, input)
	require.Nil(s.t, err, "server error")
	require.Nil(s.t, payload.Error, "client error")
	return s
}

func (s *ApiDatabaseStage) EnvGrantMutation(input model.EnvGrantInput) *ApiDatabaseStage {
	ctx, cancel := testContext()
	defer cancel()

	payload, err := s.server.EnvGrantMutation(ctx, input)
	require.Nil(s.t, err, "server error")
	require.Nil(s.t, payload.Error, "client error")

	if payload.Token != nil {
		jwt, err := api.NewJWT()
		require.Nil(s.t, err, "api jwt error")

		values, err := jwt.Validate(*payload.Token)
		assert.Nil(s.t, err, "jwt env token invalid")
		assert.NotEmpty(s.t, values, "jwt env token empty")

		s.t.Logf("Setting env token for further communication with HTTP server - claims: %+v", values)
		s.server.SetEnvToken(*payload.Token)
	}
	return s
}

func (s *ApiDatabaseStage) EnvironmentIsCreated(env string) *ApiDatabaseStage {
	ctx, cancel := testContext()
	defer cancel()

	payload, err := s.server.CreateEnvironment(ctx, model.CreateEnvironmentInput{
		MutationID:  ptr("test-db-env"),
		Env:         env,
		Description: ptr("test-dv-env"),
	})
	require.Nil(s.t, err, "create env - server error")
	require.Nil(s.t, payload.Error, "create env - client error")

	return s
}
