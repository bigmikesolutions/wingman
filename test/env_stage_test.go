package test

import (
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib" // drivers
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // drivers
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/bigmikesolutions/wingman/graphql/model"
	"github.com/bigmikesolutions/wingman/server/a10n"
	"github.com/bigmikesolutions/wingman/test/api"
)

var envCleanTables = []string{
	"wingman.environments",
}

type EnvironmentStage struct {
	t      *testing.T
	server *api.HTTPServer
	dbx    *sqlx.DB

	createEnv *model.CreateEnvironmentPayload
	env       *model.Environment
	err       error
}

func NewEnvironmentStage(t *testing.T) *EnvironmentStage {
	dbx := mustDB()
	server, err := api.New(dbx, newProviders(dbx))
	require.Nil(t, err, "api server")

	return &EnvironmentStage{
		t:      t,
		server: server,
		dbx:    dbx,
	}
}

func (s *EnvironmentStage) Close() {
	s.server.Close()

	cleanTables(s.dbx, envCleanTables)
	_ = s.dbx.Close()
}

func (s *EnvironmentStage) Given() *EnvironmentStage {
	return s
}

func (s *EnvironmentStage) When() *EnvironmentStage {
	return s
}

func (s *EnvironmentStage) Then() *EnvironmentStage {
	return s
}

func (s *EnvironmentStage) And() *EnvironmentStage {
	return s
}

func (s *EnvironmentStage) ServerIsUpAndRunning() *EnvironmentStage {
	api.AssertHeartbeat(s.t, s.server)
	return s
}

func (s *EnvironmentStage) NoClientError() *EnvironmentStage {
	assert.NoError(s.t, s.err, "query has failed")
	return s
}

func (s *EnvironmentStage) ClientErrorIs(expMsg string) *EnvironmentStage {
	require.Error(s.t, s.err, "client error expected", expMsg)
	assert.Contains(s.t, s.err.Error(), expMsg, "unexpected client error")
	return s
}

func (s *EnvironmentStage) UserIdentity(user a10n.UserIdentity) *EnvironmentStage {
	s.server.SetUser(user)
	return s
}

func (s *EnvironmentStage) EnvironmentHasBeenCreated(input model.CreateEnvironmentInput) *EnvironmentStage {
	ctx, cancel := testContext()
	defer cancel()

	payload, err := s.server.CreateEnvironment(ctx, input)
	require.Nil(s.t, err, "create env - server error")
	require.Nil(s.t, payload.Error, "create env - client error")

	return s
}

func (s *EnvironmentStage) CreateEnvironmentMutation(input model.CreateEnvironmentInput) *EnvironmentStage {
	ctx, cancel := testContext()
	defer cancel()

	s.createEnv, s.err = s.server.CreateEnvironment(ctx, input)

	return s
}

func (s *EnvironmentStage) CreateEnvironmentPayloadIs(expPayload model.CreateEnvironmentPayload) *EnvironmentStage {
	assert.NotNil(s.t, s.createEnv, "create env payload is required")
	if s.createEnv == nil {
		assert.Equal(s.t, expPayload, *s.createEnv, "unexpected create env payload")
	}

	return s
}

func (s *EnvironmentStage) EnvironmentQuery(env string) *EnvironmentStage {
	ctx, cancel := testContext()
	defer cancel()

	s.env, s.err = s.server.EnvironmentQuery(ctx, env)

	return s
}

func (s *EnvironmentStage) EnvironmentQueryResponseIs(expEnv model.Environment) *EnvironmentStage {
	assert.NotNil(s.t, s.env, "query payload is required")
	if s.env != nil {
		// skip comparison
		expEnv.CreatedAt = s.env.CreatedAt
		expEnv.ModifiedAt = s.env.ModifiedAt
		// compare
		assert.Equal(s.t, expEnv, *s.env, "unexpected env query data")
	}

	return s
}

func (s *EnvironmentStage) NoEnvironmentIsReturned() *EnvironmentStage {
	assert.Nil(s.t, s.env, "query payload must be nil")
	return s
}

func (s *EnvironmentStage) EnvGrantMutation(input model.EnvGrantInput) *EnvironmentStage {
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
