package conv

import (
	"github.com/bigmikesolutions/wingman/graphql/model"
	"github.com/bigmikesolutions/wingman/providers/db"
)

func InputToInfo(input model.AddDatabaseInput) db.ConnectionInfo {
	return db.ConnectionInfo{
		ID:     input.ID,
		Env:    input.Env,
		Host:   input.Host,
		Port:   input.Port,
		Name:   input.Name,
		User:   input.User,
		Pass:   input.Password,
		Driver: driverType(input.Driver),
	}
}

func driverType(v model.DriverType) string {
	switch v {
	case model.DriverTypePostgres:
		return db.PostgresDriverName

	default:
		return "unknown"
	}
}
