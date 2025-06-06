package containers

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

const (
	postgresNoSSL      = "sslmode=disable"
	PostgresDriverName = "pgx"
)

// Config represents configuration of docker compose with random ports.
type Config struct {
	Postgres   PostgresCfg
	ToxiProxy  ToxiProxyCfg
	Vault      VaultCfg
	Localstack LocalstackCfg
	uid        int
}

// PostgresCfg keeps docker-compose config.
type PostgresCfg struct {
	Name string
	User string
	Pass string
	Port int
}

// ToxiProxyCfg keeps docker-compose config.
type ToxiProxyCfg struct {
	Port         int
	PostgresPort int
}

// VaultCfg keeps docker-compose config.
type VaultCfg struct {
	Port      int
	RootToken string
}

// LocalstackCfg keeps docker-compose config.
type LocalstackCfg struct {
	Port   int
	Region string
}

func newCfg() Config {
	return Config{
		uid: rand.Int() % 1024,
		Postgres: PostgresCfg{
			Port: randomPort(),
			Name: "wingman",
			User: "postgres",
			Pass: "pass",
		},
		ToxiProxy: ToxiProxyCfg{
			Port:         randomPort(),
			PostgresPort: randomPort(),
		},
		Vault: VaultCfg{
			Port:      randomPort(),
			RootToken: "root",
		},
		Localstack: LocalstackCfg{
			Port:   randomPort(),
			Region: "eu-west-1",
		},
	}
}

// Env returns env variables.
func (c Config) Env() map[string]string {
	return map[string]string{
		"uid":              strconv.Itoa(c.uid),
		"pgPort":           strconv.Itoa(c.Postgres.Port),
		"pgName":           c.Postgres.Name,
		"pgUser":           c.Postgres.User,
		"pgPass":           c.Postgres.Pass,
		"toxiPort":         strconv.Itoa(c.ToxiProxy.Port),
		"toxiPgPort":       strconv.Itoa(c.ToxiProxy.PostgresPort),
		"vaultPort":        strconv.Itoa(c.Vault.Port),
		"vaultRootToken":   c.Vault.RootToken,
		"localstackPort":   strconv.Itoa(c.Localstack.Port),
		"localstackRegion": c.Localstack.Region,
	}
}

// ConnectionString returns connection string for postgres container.
func (c PostgresCfg) ConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?%s",
		c.User, c.Pass,
		GetHost(), c.Port,
		c.Name,
		postgresNoSSL,
	)
}
