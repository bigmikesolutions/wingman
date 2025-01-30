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
	secureStorage interface {
		Write(context.Context, string, any) error
		Read(context.Context, string, any) error
	}

	rbac interface {
		WriteInfo(ctx context.Context) error
		ReadInfo(ctx context.Context, dbID string) error
		ReadConnection(ctx context.Context, dbID string) error
		ReadTable(ctx context.Context, dbID string, tableName string, columns ...string) error
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
	return s.storage.Write(ctx, path(db.ID), db)
}

func (s *Service) Connection(ctx context.Context, id ID) (*Connection, error) {
	if err := s.rbac.ReadConnection(ctx, id); err != nil {
		return nil, err
	}

	conn, ok := s.conns[id]
	if !ok {
		dbInfo, err := s.Info(ctx, id)
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
			db:   dbx,
			rbac: s.rbac,
		}
		s.conns[id] = conn
	}
	return conn, nil
}

func (s *Service) Info(ctx context.Context, id ID) (*ConnectionInfo, error) {
	if err := s.rbac.ReadInfo(ctx, id); err != nil {
		return nil, err
	}

	conn := &ConnectionInfo{}
	if err := s.storage.Read(ctx, path(id), &conn); err != nil {
		return nil, err
	}
	return conn, nil
}

func (s *Service) Close() error {
	return s.rbac.Close()
}

func path(id ID) string {
	return fmt.Sprintf("/providers/db/connections/%s", id)
}
