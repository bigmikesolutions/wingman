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
	"github.com/bigmikesolutions/wingman/server/a10n"
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
	"wingman.environments",
	"provider_db.user_role",
}

type DatabaseStage struct {
	t      *testing.T
	server *api.HTTPServer
	dbx    *sqlx.DB

	queryDatabase *model.Database
	err           error
}

func NewDatabaseStage(t *testing.T) *DatabaseStage {
	dbx := mustDB()
	server, err := api.New(dbx, newProviders(dbx))
	require.Nil(t, err, "api server")

	return &DatabaseStage{
		t:      t,
		server: server,
		dbx:    dbx,
	}
}

func (s *DatabaseStage) Close() {
	s.server.Close()

	cleanTables(s.dbx, dbCleanTables)
	_ = s.dbx.Close()
}

func (s *DatabaseStage) Given() *DatabaseStage {
	return s
}

func (s *DatabaseStage) When() *DatabaseStage {
	return s
}

func (s *DatabaseStage) Then() *DatabaseStage {
	return s
}

func (s *DatabaseStage) And() *DatabaseStage {
	return s
}

func (s *DatabaseStage) ServerIsUpAndRunning() *DatabaseStage {
	api.AssertHeartbeat(s.t, s.server)
	return s
}

func (s *DatabaseStage) DatabaseInfoQuery(
	env string,
	id string,
) *DatabaseStage {
	ctx, cancel := testContext()
	defer cancel()
	s.queryDatabase, s.err = s.server.DatabaseInfoQuery(ctx, env, id)
	return s
}

func (s *DatabaseStage) QueryDatabaseTableData(
	env string,
	id string,
	name string,
	first int,
	after cursor.Cursor,
	where model.TableFilter,
) *DatabaseStage {
	ctx, cancel := testContext()
	defer cancel()
	s.queryDatabase, s.err = s.server.DatabaseTableDataQuery(ctx, env, id, name, first, after, where)
	return s
}

func (s *DatabaseStage) NoClientError() *DatabaseStage {
	assert.NoError(s.t, s.err, "query has failed")
	return s
}

func (s *DatabaseStage) ClientErrorIs(expMsg string) *DatabaseStage {
	require.Error(s.t, s.err, "client error expected", expMsg)
	assert.Contains(s.t, s.err.Error(), expMsg, "unexpected client error")
	return s
}

func (s *DatabaseStage) DatabaseIsProvided(database db.ConnectionInfo) *DatabaseStage {
	ctx, cancel := testContext()
	defer cancel()
	err := s.server.Resolver.Providers.DB.Register(ctx, database)
	require.Nilf(s.t, err, "register database: %+v", database)
	return s
}

func (s *DatabaseStage) DatabaseStatement(dbID, statement string, args ...any) *DatabaseStage {
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

func (s *DatabaseStage) DatabaseInfoIsReturned(dbID, driver string) *DatabaseStage {
	require.NotNil(s.t, s.queryDatabase, "query database is nil")

	assert.Equal(s.t, dbID, s.queryDatabase.ID, "database ID")
	assert.Equal(s.t, dbID, s.queryDatabase.Info.ID, "database info: ID")
	assert.Equal(s.t, driver, string(s.queryDatabase.Info.Driver), "database info: driver")
	assert.Equal(s.t, "localhost", s.queryDatabase.Info.Host, "database info: host")
	assert.NotEqual(s.t, 0, s.queryDatabase.Info.Port, "database info: port")
	return s
}

func (s *DatabaseStage) TableQueryHasNextPage(nextPage bool) *DatabaseStage {
	require.NotNil(s.t, s.queryDatabase.Table, "table data must be returned")
	require.NotNil(s.t, s.queryDatabase.Table.ConnectionInfo, "connection info must be returned")

	assert.Equal(s.t, nextPage, s.queryDatabase.Table.ConnectionInfo.HasNextPage, "next page")
	if nextPage {
		assert.NotEqual(s.t, "", string(s.queryDatabase.Table.ConnectionInfo.EndCursor), "end cursor missing")
	}
	return s
}

func (s *DatabaseStage) TableHasRows(expRows ...model.TableRow) *DatabaseStage {
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

func (s *DatabaseStage) NoTableRows() *DatabaseStage {
	require.NotNil(s.t, s.queryDatabase.Table, "table data must be returned")

	for _, edge := range s.queryDatabase.Table.Edges {
		assert.Empty(s.t, edge.Node.Values, "no table rows expected")
	}

	return s
}

func (s *DatabaseStage) DatabaseUserRoleHasBeenCreated(input model.AddDatabaseUserRoleInput) *DatabaseStage {
	ctx, cancel := testContext()
	defer cancel()
	payload, err := s.server.CreateDatabaseUserRoleMutation(ctx, input)
	require.Nil(s.t, err, "server error")
	require.Nil(s.t, payload.Error, "client error")
	return s
}

func (s *DatabaseStage) EnvGrantMutation(input model.EnvGrantInput) *DatabaseStage {
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

func (s *DatabaseStage) EnvironmentHasBeenCreated(env string) *DatabaseStage {
	ctx, cancel := testContext()
	defer cancel()

	payload, err := s.server.CreateEnvironment(ctx, model.CreateEnvironmentInput{
		MutationID:  ptr(s.t.Name()),
		Env:         env,
		Description: ptr(s.t.Name()),
	})
	require.Nil(s.t, err, "create env - server error")
	require.Nil(s.t, payload.Error, "create env - client error")

	return s
}

func (s *DatabaseStage) UserIdentity(user a10n.UserIdentity) *DatabaseStage {
	s.server.SetUser(user)
	return s
}
