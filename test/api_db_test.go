package test

import (
	"testing"

	"github.com/bigmikesolutions/wingman/graphql/model"
)

func Test_Api_Database_ShouldGetInfo(t *testing.T) {
	s := NewApiDatabaseStage(t)
	defer s.Close()

	s.Given().
		ServerIsUpAndRunning().And().
		DatabaseIsProvided(connectionInfo("pg-1", dc.Postgres()))

	s.When().
		QueryDatabase("test-env", "pg-1")

	s.Then().
		NoClientError().And().
		DatabaseInfoIsReturned("pg-1", "POSTGRES")
}

func Test_Api_Database_ShouldGetTableData(t *testing.T) {
	s := NewApiDatabaseStage(t)
	defer s.Close()

	s.Given().
		ServerIsUpAndRunning().And().
		DatabaseIsProvided(connectionInfo("pg-1", dc.Postgres())).
		DatabaseStatement("pg-1", sqlCreateTableStudents).And().
		DatabaseStatement("pg-1", sqlInsertStudents)

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
		NoClientError().And().
		DatabaseInfoIsReturned("pg-1", "POSTGRES").And().
		TableQueryHasNextPage(false).And().
		TableHasRows(
			model.TableRow{
				Index:  ptr(0),
				Values: []*string{ptr("1"), ptr("johny"), ptr("bravo"), ptr("30")},
			},
			model.TableRow{
				Index:  ptr(1),
				Values: []*string{ptr("2"), ptr("mike"), ptr("tyson"), ptr("51")},
			},
			model.TableRow{
				Index:  ptr(2),
				Values: []*string{ptr("3"), ptr("pamela"), ptr("anderson"), ptr("65")},
			},
		)
}
