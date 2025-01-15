// Package containers wraps-up containers set-up.
package containers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	toxiproxy "github.com/Shopify/toxiproxy/v2/client"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	dockerID = "wingman"
)

// Service holds state of started containers.
type Service struct {
	cfg          cfg
	composeStack compose.ComposeStack
	toxiClient   *toxiproxy.Client
}

// New start containers.
func New(ctx context.Context) (*Service, error) {
	cfg := newCfg()

	composeStack, err := compose.NewDockerComposeWith(
		compose.WithStackFiles("containers/docker-compose.yml"),
		compose.WithLogger(&logger{}),
	)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		cfg:          cfg,
		composeStack: composeStack,
		toxiClient:   toxiproxy.NewClient(fmt.Sprintf("%s:%d", GetHost(), cfg.ToxiProxy.Port)),
	}

	err = composeStack.
		WithEnv(cfg.Env()).
		WaitForService("postgres", wait.ForListeningPort("5432/tcp")).
		Up(ctx, compose.Wait(true))
	if err != nil {
		defer svc.Close()
		return nil, fmt.Errorf("docker-compose up: %w", err)
	}

	return svc, nil
}

// Postgres returns container.
func (s *Service) Postgres() PostgresCfg {
	return s.cfg.Postgres
}

// DB creates direct connect to database.
func (s *Service) DB(ctx context.Context) (*sqlx.DB, error) {
	return sqlx.ConnectContext(ctx, PostgresDriverName, s.cfg.Postgres.ConnectionString())
}

// DBProxy creates connection to database via proxy.
func (s *Service) DBProxy(ctx context.Context) (*DBProxy, error) {
	containerName := fmt.Sprintf("%s_postgres_%s", dockerID, s.cfg.uid)
	upstream := fmt.Sprintf("%s:%d", containerName, s.cfg.Postgres.Port)
	listen := fmt.Sprintf("[::]:%d", s.cfg.ToxiProxy.PostgresPort)

	proxy, err := s.toxiClient.CreateProxy(containerName, listen, upstream)

	connString := s.cfg.Postgres.ConnectionString()
	connString = strings.ReplaceAll(connString, strconv.Itoa(s.cfg.Postgres.Port), strconv.Itoa(s.cfg.ToxiProxy.PostgresPort))

	conn, err := sqlx.ConnectContext(ctx, PostgresDriverName, connString)
	if err != nil {
		return nil, fmt.Errorf("postgres proxy connection: %w", err)
	}

	return &DBProxy{
		DB: conn,
		Upstream: &Upstream{
			proxy: proxy,
		},
	}, nil
}

// Close closes all created containers gracefully.
func (s *Service) Close() {
	if err := s.composeStack.Down(
		context.Background(),
		compose.RemoveOrphans(true),
		compose.RemoveImagesLocal,
	); err != nil {
		log.Printf("docker-compose down: %v", err)
	}
}
