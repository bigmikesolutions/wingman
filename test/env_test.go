package test

import (
	"testing"

	"github.com/bigmikesolutions/wingman/graphql/model"
	"github.com/bigmikesolutions/wingman/server/a10n"
)

func Test_Environment_ShouldCreateEnvironment(t *testing.T) {
	s := NewEnvironmentStage(t)
	defer s.Close()

	envID := "test-env"

	s.Given().
		ServerIsUpAndRunning().And().
		UserIdentity(a10n.UserIdentity{
			UserID: testUserID,
			OrgID:  testOrg,
			Roles:  []a10n.Role{a10n.AdminWrite, a10n.AdminRead},
		})

	s.When().
		CreateEnvironmentMutation(model.CreateEnvironmentInput{
			MutationID: ptr(t.Name()),
			Env:        envID,
		})

	s.Then().
		NoClientError().And().
		CreateEnvironmentPayloadIs(model.CreateEnvironmentPayload{
			MutationID: ptr(t.Name()),
			Env:        envID,
		})
}

func Test_Environment_ShouldNotDuplicateEnvironmentWithinOrg(t *testing.T) {
	s := NewEnvironmentStage(t)
	defer s.Close()

	envID := "test-env"

	s.Given().
		ServerIsUpAndRunning().And().
		UserIdentity(a10n.UserIdentity{
			UserID: testUserID,
			OrgID:  testOrg,
			Roles:  []a10n.Role{a10n.AdminWrite, a10n.AdminRead},
		}).And().
		EnvironmentHasBeenCreated(model.CreateEnvironmentInput{
			MutationID: ptr(t.Name()),
			Env:        envID,
		})

	s.When().
		CreateEnvironmentMutation(model.CreateEnvironmentInput{
			MutationID: ptr(t.Name()),
			Env:        envID,
		})

	s.Then().
		ClientErrorIs("environment already exists")
}

func Test_Environment_ShouldDuplicateEnvironmentForDifferentOrganisations(t *testing.T) {
	s := NewEnvironmentStage(t)
	defer s.Close()

	envID := "test-env"

	s.Given().
		ServerIsUpAndRunning().And().
		UserIdentity(a10n.UserIdentity{
			UserID: testUserID,
			OrgID:  testOrg,
			Roles:  []a10n.Role{a10n.AdminWrite, a10n.AdminRead},
		}).And().
		EnvironmentHasBeenCreated(model.CreateEnvironmentInput{
			MutationID: ptr(t.Name()),
			Env:        envID,
		})

	s.When().
		UserIdentity(a10n.UserIdentity{
			UserID: "diff-user",
			OrgID:  "diff-org",
			Roles:  []a10n.Role{a10n.AdminWrite, a10n.AdminRead},
		}).And().
		CreateEnvironmentMutation(model.CreateEnvironmentInput{
			MutationID: ptr(t.Name()),
			Env:        envID,
		})

	s.Then().
		NoClientError().And().
		CreateEnvironmentPayloadIs(model.CreateEnvironmentPayload{
			MutationID: ptr(t.Name()),
			Env:        envID,
		})
}

func Test_Environment_ShouldNotCreateEnvironmentForInsufficientPrivileges(t *testing.T) {
	s := NewEnvironmentStage(t)
	defer s.Close()

	envID := "test-env"

	s.Given().
		ServerIsUpAndRunning().And().
		UserIdentity(a10n.UserIdentity{
			UserID: testUserID,
			OrgID:  testOrg,
			Roles:  []a10n.Role{a10n.AdminRead},
		})

	s.When().
		CreateEnvironmentMutation(model.CreateEnvironmentInput{
			MutationID: ptr(t.Name()),
			Env:        envID,
		})

	s.Then().
		ClientErrorIs("user not authorized")
}

func Test_Environment_ShouldReadExistingEnvironment(t *testing.T) {
	s := NewEnvironmentStage(t)
	defer s.Close()

	envID := "test-env"

	s.Given().
		ServerIsUpAndRunning().And().
		UserIdentity(a10n.UserIdentity{
			UserID: testUserID,
			OrgID:  testOrg,
			Roles:  []a10n.Role{a10n.AdminWrite, a10n.AdminRead},
		}).And().
		EnvironmentHasBeenCreated(model.CreateEnvironmentInput{
			MutationID:  ptr(t.Name()),
			Env:         envID,
			Description: ptr(t.Name()),
		}).And().
		EnvGrantMutation(model.EnvGrantInput{
			MutationID: ptr(t.Name()),
			Reason:     ptr(t.Name()),
			Resource: []*model.ResourceGrantInput{
				{
					Env: envID,
				},
			},
		})

	s.When().
		EnvironmentQuery(envID)

	s.Then().
		NoClientError().And().
		EnvironmentQueryResponseIs(model.Environment{
			ID:          envID,
			Description: ptr(t.Name()),
		})
}

func Test_Environment_ShouldReadNonExistingEnvironment(t *testing.T) {
	s := NewEnvironmentStage(t)
	defer s.Close()

	envID := "test-env"

	s.Given().
		ServerIsUpAndRunning().And().
		UserIdentity(a10n.UserIdentity{
			UserID: testUserID,
			OrgID:  testOrg,
			Roles:  []a10n.Role{a10n.AdminWrite, a10n.AdminRead},
		}).And().
		EnvGrantMutation(model.EnvGrantInput{
			MutationID: ptr(t.Name()),
			Reason:     ptr(t.Name()),
			Resource: []*model.ResourceGrantInput{
				{
					Env: envID,
				},
			},
		})

	s.When().
		EnvironmentQuery(envID)

	s.Then().
		NoClientError().And().
		NoEnvironmentIsReturned()
}

func Test_Environment_ShouldNotReadEnvironmentOfOtherOrganisation(t *testing.T) {
	s := NewEnvironmentStage(t)
	defer s.Close()

	envID := "test-env"

	s.Given().
		ServerIsUpAndRunning().And().
		UserIdentity(a10n.UserIdentity{
			UserID: testUserID,
			OrgID:  testOrg,
			Roles:  []a10n.Role{a10n.AdminWrite, a10n.AdminRead},
		}).And().
		EnvironmentHasBeenCreated(model.CreateEnvironmentInput{
			MutationID:  ptr(t.Name()),
			Env:         envID,
			Description: ptr(t.Name()),
		}).And().
		EnvGrantMutation(model.EnvGrantInput{
			MutationID: ptr(t.Name()),
			Reason:     ptr(t.Name()),
			Resource: []*model.ResourceGrantInput{
				{
					Env: envID,
				},
			},
		})

	s.When().
		UserIdentity(a10n.UserIdentity{
			UserID: "diff user",
			OrgID:  "diff org",
			Roles:  []a10n.Role{a10n.AdminWrite, a10n.AdminRead},
		}).And().
		EnvironmentQuery(envID)

	s.Then().
		NoClientError().And().
		NoEnvironmentIsReturned()
}

func Test_Environment_ShouldNotReadEnvironmentIfNoAccessGranted(t *testing.T) {
	s := NewEnvironmentStage(t)
	defer s.Close()

	envID := "test-env"

	s.Given().
		ServerIsUpAndRunning().And().
		UserIdentity(a10n.UserIdentity{
			UserID: testUserID,
			OrgID:  testOrg,
			Roles:  []a10n.Role{a10n.AdminWrite, a10n.AdminRead},
		}).And().
		EnvironmentHasBeenCreated(model.CreateEnvironmentInput{
			MutationID:  ptr(t.Name()),
			Env:         envID,
			Description: ptr(t.Name()),
		})

	s.When().
		EnvironmentQuery(envID)

	s.Then().
		ClientErrorIs("no active environment session").And().
		NoEnvironmentIsReturned()
}

func Test_Environment_ShouldNotReadEnvironmentForInsufficientPrivileges(t *testing.T) {
	s := NewEnvironmentStage(t)
	defer s.Close()

	envID := "test-env"

	s.Given().
		ServerIsUpAndRunning().And().
		UserIdentity(a10n.UserIdentity{
			UserID: testUserID,
			OrgID:  testOrg,
			Roles:  []a10n.Role{a10n.AdminWrite, a10n.AdminRead},
		}).And().
		EnvironmentHasBeenCreated(model.CreateEnvironmentInput{
			MutationID:  ptr(t.Name()),
			Env:         envID,
			Description: ptr(t.Name()),
		}).And().
		EnvGrantMutation(model.EnvGrantInput{
			MutationID: ptr(t.Name()),
			Reason:     ptr(t.Name()),
			Resource: []*model.ResourceGrantInput{
				{
					Env: envID,
				},
			},
		})

	s.When().
		UserIdentity(a10n.UserIdentity{
			UserID: testUserID,
			OrgID:  testOrg,
			Roles:  []a10n.Role{},
		}).And().
		EnvironmentQuery(envID)

	s.Then().
		ClientErrorIs("user not authorized").And().
		NoEnvironmentIsReturned()
}
