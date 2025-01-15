package containers

import (
	"fmt"
	"math/rand"
	"strconv"
)

const (
	postgresNoSSL      = "sslmode=disable"
	PostgresDriverName = "pgx"
)

// cfg represents configuration of docker compose with random ports
type cfg struct {
	uid       int
	Postgres  PostgresCfg
	ToxiProxy ToxiProxyCfg
}

type PostgresCfg struct {
	Port int
	Name string
	User string
	Pass string
}

type ToxiProxyCfg struct {
	Port         int
	PostgresPort int
}

func newCfg() cfg {
	return cfg{
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
	}
}

func (c cfg) Env() map[string]string {
	return map[string]string{
		"uid":        strconv.Itoa(c.uid),
		"pgPort":     strconv.Itoa(c.Postgres.Port),
		"pgName":     c.Postgres.Name,
		"pgUser":     c.Postgres.User,
		"pgPass":     c.Postgres.Pass,
		"toxiPort":   strconv.Itoa(c.ToxiProxy.Port),
		"toxiPgPort": strconv.Itoa(c.ToxiProxy.PostgresPort),
	}
}

func (c PostgresCfg) ConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?%s",
		c.User, c.Pass,
		GetHost(), c.Port,
		c.Name,
		postgresNoSSL,
	)
}
