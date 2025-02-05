package test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/bigmikesolutions/wingman/graphql/model"
)

func Test_Api_Database_ShouldGrantAccessToEnv(t *testing.T) {
	s := NewApiDatabaseStage(t)
	defer s.Close()

	envID := "test-env"
	dbID := uuid.New().String()
	dbCfg := newTestDatabase(dc)

	s.Given().
		ServerIsUpAndRunning().And().
		DatabaseIsProvided(connectionInfo(dbID, dbCfg))

	s.When().
		EnvGrantMutation(model.EnvGrantInput{
			MutationID: ptr("creat-env-grant"),
			Reason:     ptr("testing..."),
			Resource: []*model.ResourceGrantInput{
				{
					Env: "test-env",
					Database: []*model.DatabaseResource{
						{
							ID:    envID,
							Table: []*model.TableResource{},
						},
					},
				},
			},
		})

	s.Then().
		NoClientError()
}

func Test_Api_Database_ShouldGetInfo(t *testing.T) {
	s := NewApiDatabaseStage(t)
	defer s.Close()

	envID := "test-env"
	dbID := uuid.New().String()
	dbCfg := newTestDatabase(dc)

	s.Given().
		ServerIsUpAndRunning().And().
		DatabaseIsProvided(connectionInfo(dbID, dbCfg))

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
	roleID := uuid.New().String()

	dbID := uuid.New().String()
	dbCfg := newTestDatabase(dc)

	s.Given().
		ServerIsUpAndRunning().And().
		DatabaseIsProvided(connectionInfo(dbID, dbCfg)).
		DatabaseStatement(dbID, sqlCreateTableStudents).And().
		DatabaseStatement(dbID, sqlInsertStudents).And().
		DatabaseUserRoleIsCreated(model.AddDatabaseUserRoleInput{
			MutationID: ptr(t.Name()),
			UserRoles: []*model.AddDatabaseUserRole{
				{
					ID:          ptr(roleID),
					Description: ptr("read-only access"),
					DatabaseAccess: []*model.DatabaseAccessInput{
						{
							ID: dbID,
							Tables: []*model.DatabaseTableAccessInput{
								{
									Name:       sqlTableStudents,
									AccessType: model.AccessTypeReadOnly,
								},
							},
						},
					},
				},
			},
		})

	s.When().
		QueryDatabaseTableData(
			envID,
			dbID,
			sqlTableStudents,
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

func Test_Api_Database_ShouldForbidGettingTableData(t *testing.T) {
	s := NewApiDatabaseStage(t)
	defer s.Close()

	envID := "test-env"
	roleID := uuid.New().String()

	dbID := uuid.New().String()
	dbCfg := newTestDatabase(dc)

	s.Given().
		ServerIsUpAndRunning().And().
		DatabaseIsProvided(connectionInfo(dbID, dbCfg)).
		DatabaseStatement(dbID, sqlCreateTableStudents).And().
		DatabaseStatement(dbID, sqlInsertStudents).And().
		DatabaseUserRoleIsCreated(model.AddDatabaseUserRoleInput{
			MutationID: ptr(t.Name()),
			UserRoles: []*model.AddDatabaseUserRole{
				{
					ID:          ptr(roleID),
					Description: ptr("read-only access"),
					DatabaseAccess: []*model.DatabaseAccessInput{
						{
							ID: dbID,
							Tables: []*model.DatabaseTableAccessInput{
								{
									Name:       sqlTableNonExistingTable,
									AccessType: model.AccessTypeReadOnly,
								},
							},
						},
					},
				},
			},
		})

	s.When().
		QueryDatabaseTableData(
			envID,
			dbID,
			sqlTableStudents,
			50,
			"",
			model.TableFilter{},
		)

	s.Then().
		ClientErrorIs("database access denied") // TODO make client errors user friendly
}
