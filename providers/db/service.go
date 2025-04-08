package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/bigmikesolutions/wingman/server/a10n"

	_ "github.com/jackc/pgx/v5/stdlib" // drivers
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // drivers
)

var ErrDatabaseNotFound = errors.New("database not found")

type (
	secureStorage interface {
		Write(context.Context, string, any) error
		Read(context.Context, string, any) error
	}

	rbac interface {
		WriteInfo(ctx context.Context) error
		ReadInfo(ctx context.Context, env string, dbID string) error
		ReadConnection(ctx context.Context, env string, dbID string) error
		ReadTable(ctx context.Context, env string, dbID string, tableName string, columns ...string) error
		Close() error
	}

	Service struct {
		storage secureStorage
		conns   map[ID]*Connection
		rbac    rbac
	}
)

func New(rbac rbac, storage secureStorage) *Service {
	return &Service{
		rbac:    rbac,
		storage: storage,
		conns:   make(map[string]*Connection),
	}
}

func (s *Service) Register(ctx context.Context, db ConnectionInfo) error {
	if err := s.rbac.WriteInfo(ctx); err != nil {
		return err
	}

	user, err := a10n.GetIdentity(ctx)
	if err != nil {
		return err
	}

	return s.storage.Write(ctx, path(user.OrgID, db.Env, db.ID), db)
}

func (s *Service) Connection(ctx context.Context, env string, id ID) (*Connection, error) {
	if err := s.rbac.ReadConnection(ctx, env, id); err != nil {
		return nil, err
	}

	conn, ok := s.conns[id]
	if !ok {
		dbInfo, err := s.Info(ctx, env, id)
		if err != nil {
			// TODO handle not found error
			return nil, err
		}
		if dbInfo.ID == "" {
			return nil, ErrDatabaseNotFound
		}

		dbx, err := sqlx.ConnectContext(ctx, dbInfo.Driver, connectionString(*dbInfo))
		if err != nil {
			return nil, fmt.Errorf("connect to database: %w", err)
		}

		conn = &Connection{
			dbID: id,
			env:  env,
			db:   dbx,
			rbac: s.rbac,
		}
		s.conns[id] = conn
	}
	return conn, nil
}

func (s *Service) Info(ctx context.Context, env string, id ID) (*ConnectionInfo, error) {
	if err := s.rbac.ReadInfo(ctx, env, id); err != nil {
		return nil, err
	}

	user, err := a10n.GetIdentity(ctx)
	if err != nil {
		return nil, err
	}

	conn := &ConnectionInfo{}
	if err := s.storage.Read(ctx, path(user.OrgID, env, id), &conn); err != nil {
		return nil, err
	}
	return conn, nil
}

func (s *Service) Close() error {
	return s.rbac.Close()
}

func path(orgID, env string, id ID) string {
	return fmt.Sprintf(
		"/providers/db/organisations/%s/environments/%s/connections/%s",
		orgID, env, id,
	)
}
