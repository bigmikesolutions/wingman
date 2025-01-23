package db

import (
	"context"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // drivers
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // drivers
)

var ErrDatabaseNotFound = errors.New("database not found")

type (
	Service struct {
		storage secureStorage
		conns   map[ID]*Connection
		rbac    RBAC
	}
)

func New(rbac RBAC, storage secureStorage) *Service {
	return &Service{
		rbac:    rbac,
		storage: storage,
		conns:   make(map[string]*Connection),
	}
}

func (s *Service) Register(ctx context.Context, db ConnectionInfo) error {
	return s.storage.write(ctx, db)
}

func (s *Service) Connection(ctx context.Context, id ID) (*Connection, error) {
	conn, ok := s.conns[id]
	if !ok {
		dbInfo, err := s.storage.read(ctx, id)
		if err != nil {
			return nil, err
		}
		if dbInfo == nil {
			return nil, ErrDatabaseNotFound
		}

		roles, rolesErr := s.rbac.FindUserRolesByDatabaseID(ctx, id) // TODO use user ID from context here
		if rolesErr != nil {
			return nil, fmt.Errorf("check database roles: %w", rolesErr)
		}
		if len(roles) == 0 {
			return nil, ErrDatabaseAccessDenied
		}

		dbx, err := sqlx.ConnectContext(ctx, dbInfo.Driver, connectionString(*dbInfo))
		if err != nil {
			return nil, fmt.Errorf("connect to database: %w", err)
		}

		conn = &Connection{
			dbID: id,
			db:   dbx,
			rbac: s.rbac,
		}
		s.conns[id] = conn
	}
	return conn, nil
}

func (s *Service) Info(ctx context.Context, id ID) (*ConnectionInfo, error) {
	return s.storage.read(ctx, id)
}

func (s *Service) RBAC() RBAC {
	// TODO check user access here
	return s.rbac
}

func (s *Service) Close() error {
	return s.rbac.Close()
}
