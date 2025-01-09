package test

import (
	"github.com/google/uuid"
	"testing"

	"github.com/bigmikesolutions/wingman/graphql/model"
)

func Test_Api_Database_ShouldGetInfo(t *testing.T) {
	s := NewApiDatabaseStage(t)
	defer s.Close()

	envID := "test-env"
	dbID := "pg-1"

	s.Given().
		ServerIsUpAndRunning().And().
		DatabaseIsProvided(connectionInfo(dbID, dc.Postgres()))

	s.When().
		QueryDatabase(envID, dbID)

	s.Then().
		NoClientError().And().
		DatabaseInfoIsReturned(dbID, expDriverPostgres)
}

func Test_Api_Database_ShouldGetTableData(t *testing.T) {
	s := NewApiDatabaseStage(t)
	defer s.Close()

	envID := "test-env"
	dbID := "pg-1"
	roleID := uuid.New().String()

	s.Given().
		ServerIsUpAndRunning().And().
		DatabaseIsProvided(connectionInfo(dbID, dc.Postgres())).
		DatabaseStatement(dbID, sqlCreateTableStudents).And().
		DatabaseStatement(dbID, sqlInsertStudents).And().
		DatabaseUserRoleIsCreated(model.AddDatabaseUserRoleInput{
			MutationID: ptr(t.Name()),
			UserRoles: []*model.AddDatabaseUserRole{
				{
					ID:          ptr(roleID),
					AccessType:  model.AccessTypeRead,
					Description: ptr("read-only access"),
					DatabaseIds: []*string{
						ptr(dbID),
					},
				},
			},
		})

	s.When().
		QueryDatabaseTableData(
			envID,
			dbID,
			"students",
			50,
			"",
			model.TableFilter{},
		)

	s.Then().
		NoClientError().And().
		DatabaseInfoIsReturned(dbID, expDriverPostgres).And().
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
