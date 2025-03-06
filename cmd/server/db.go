package main

import (
	"fmt"
	"net/url"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // drivers
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // drivers
)

type DatabaseConfig struct {
	Host                  string        `envconfig:"DB_HOST" default:"localhost"`
	Port                  int           `envconfig:"DB_PORT" default:"5432"`
	DriverName            string        `envconfig:"DB_DRIVER" default:"pgx"`
	SSLMode               string        `envconfig:"DB_SSL_MODE" default:"disable"`
	SSLRootCert           string        `envconfig:"DB_SSL_ROOT_CERT" default:""`
	DatabaseName          string        `envconfig:"DB_DATABASE" default:"bsp"`
	DatabaseSchema        string        `envconfig:"DB_SCHEMA" default:""`
	Username              string        `envconfig:"DB_USERNAME"`
	Password              string        `envconfig:"DB_PASSWORD"`
	MaxOpenConnections    int           `envconfig:"DB_MAX_OPEN_CONN" default:"10"`
	MaxIdleConnections    int           `envconfig:"DB_MAX_IDLE_CONN" default:"10"`
	ConnectionMaxLifeTime time.Duration `envconfig:"DB_CONN_MAX_LIFE_TIME" default:"5m"`
	ConnectionMaxIdleTime time.Duration `envconfig:"DB_CONN_MAX_IDLE_TIME" default:"5m"`
}

func newDB(cfg DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(cfg.DriverName, cfg.ConnectionString())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.MaxIdleConnections)
	db.SetConnMaxLifetime(cfg.ConnectionMaxLifeTime)
	db.SetConnMaxIdleTime(cfg.ConnectionMaxIdleTime)

	return db, nil
}

func (c DatabaseConfig) ConnectionString() string {
	sslRootCert := ""
	if c.SSLRootCert != "" {
		sslRootCert = fmt.Sprintf("&sslrootcert=%s", c.SSLRootCert)
	}

	searchPath := ""
	if c.DatabaseSchema != "" {
		searchPath = fmt.Sprintf("&search_path=%s", c.DatabaseSchema)
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s%s%s",
		c.Username,
		url.QueryEscape(c.Password),
		c.Host,
		c.Port,
		c.DatabaseName,
		c.SSLMode,
		sslRootCert,
		searchPath,
	)
}
