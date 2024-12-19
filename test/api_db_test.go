package test

import (
	"testing"

	"github.com/bigmikesolutions/wingman/graphql/model"
	"github.com/bigmikesolutions/wingman/providers/db"
)

func Test_Api_Database_ShouldGetInfo(t *testing.T) {
	s := NewApiDatabaseStage(t)
	defer s.Close()

	s.Given().
		ServerIsUpAndRunning().And().
		DatabaseIsProvided(db.ConnectionInfo{
			ID:     "pg-1",
			Driver: driverPGX,
			Host:   "localhost",
			Name:   "postgres",
			Port:   5432,
			User:   "admin",
			Pass:   "some-pass",
		})

	s.When().
		QueryDatabase("test-env", "pg-1")

	s.Then().
		NoClientError()
}

func Test_Api_Database_ShouldGetTableData(t *testing.T) {
	s := NewApiDatabaseStage(t)
	defer s.Close()

	s.Given().
		ServerIsUpAndRunning().And().
		DatabaseIsProvided(db.ConnectionInfo{
			ID:     "pg-1",
			Driver: driverPGX,
			Host:   "localhost",
			Name:   "postgres",
			Port:   5432,
			User:   "admin",
			Pass:   "some-pass",
		})

	s.When().
		QueryDatabaseTableData(
			"test-env",
			"pg-1",
			"students",
			50,
			"",
			model.TableFilter{},
		)

	s.Then().
		NoClientError()
}
