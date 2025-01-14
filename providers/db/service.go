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

	ConnectionInfo struct {
		ID     ID
		Driver string
		Host   string
		Name   string
		Port   int
		User   string
		Pass   string
	}

	Service struct {
		databases   map[ID]ConnectionInfo
		connections map[ID]*sqlx.DB
	}
)

func New() *Service {
	return &Service{
		connections: make(map[string]*sqlx.DB),
		databases:   make(map[string]ConnectionInfo),
	}
}

func (s *Service) Register(db ConnectionInfo) error {
	if _, ok := s.databases[db.ID]; ok {
		return ErrDatabaseAlreadyExists
	}
	s.databases[db.ID] = db
	return nil
}

func (s *Service) Connection(ctx context.Context, id ID) (*sqlx.DB, error) {
	conn, ok := s.connections[id]
	if !ok {
		db, dbOK := s.databases[id]
		if !dbOK {
			return nil, ErrDatabaseNotFound
		}

		var err error
		conn, err = sqlx.ConnectContext(ctx, db.Driver, connectionString(db))
		if err != nil {
			return nil, fmt.Errorf("connect to database: %w", err)
		}
		s.connections[id] = conn
	}
	return conn, nil
}

func (s *Service) Info(id ID) (ConnectionInfo, bool) {
	info, ok := s.databases[id]
	return info, ok
}

func connectionString(db ConnectionInfo) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		db.User, db.Pass,
		db.Host, db.Port,
		db.Name,
	)
}
