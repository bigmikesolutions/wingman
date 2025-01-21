package db

import (
	"context"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib" // drivers
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // drivers
)

var (
	ErrDatabaseAlreadyExists = errors.New("database already exists")
	ErrDatabaseNotFound      = errors.New("database not found")
)

type (
	ID = string

	Service struct {
		dbInfo map[ID]ConnectionInfo
		conns  map[ID]*Connection
		rbac   RBAC
	}
)

func New(rbac RBAC) *Service {
	return &Service{
		conns:  make(map[string]*Connection),
		dbInfo: make(map[string]ConnectionInfo),
		rbac:   rbac,
	}
}

func (s *Service) Register(db ConnectionInfo) error {
	if _, ok := s.dbInfo[db.ID]; ok {
		return ErrDatabaseAlreadyExists
	}
	s.dbInfo[db.ID] = db
	return nil
}

func (s *Service) Connection(ctx context.Context, id ID) (*Connection, error) {
	conn, ok := s.conns[id]
	if !ok {
		dbInfo, dbOK := s.dbInfo[id]
		if !dbOK {
			return nil, ErrDatabaseNotFound
		}

		roles, rolesErr := s.rbac.FindUserRolesByDatabaseID(ctx, id) // TODO use user ID from context here
		if rolesErr != nil {
			return nil, fmt.Errorf("check database roles: %w", rolesErr)
		}
		if len(roles) == 0 {
			return nil, ErrDatabaseAccessDenied
		}

		var err error
		dbx, err := sqlx.ConnectContext(ctx, dbInfo.Driver, connectionString(dbInfo))
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

func (s *Service) Info(id ID) (ConnectionInfo, bool) {
	info, ok := s.dbInfo[id]
	return info, ok
}

func (s *Service) RBAC() RBAC {
	// TODO check user access here
	return s.rbac
}

func (s *Service) Close() error {
	return s.rbac.Close()
}
