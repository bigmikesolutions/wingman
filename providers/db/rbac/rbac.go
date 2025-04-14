package rbac

import (
	"context"
	"fmt"
	"time"

	"github.com/bigmikesolutions/wingman/server/a10n"
)

type (
	userRoleRepo interface {
		CreateUserRole(ctx context.Context, role UserRole) error
		FindUserRolesByDatabaseID(ctx context.Context, orgID string, env string, id string) ([]UserRole, error)
		Close() error
	}

	Service struct {
		repo userRoleRepo
	}
)

func New(repo userRoleRepo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUserRole(ctx context.Context, role UserRole) error {
	user, err := a10n.GetIdentity(ctx)
	if err != nil {
		return err
	}

	if err := user.ContainsRole(a10n.AdminWrite); err != nil {
		return err
	}

	role.OrgID = user.OrgID
	role.CreatedBy = user.UserID
	role.CreatedAt = time.Now()

	return s.repo.CreateUserRole(ctx, role)
}

func (s *Service) ReadInfo(ctx context.Context, env string, dbID string) error {
	user, err := a10n.GetIdentity(ctx)
	if err != nil {
		return err
	}

	roles, err := s.repo.FindUserRolesByDatabaseID(ctx, user.OrgID, env, dbID)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrDatabaseAccessDenied, err.Error())
	}

	canRead := false
	for _, role := range roles {
		if role.Info != nil && *role.Info == ReadOnlyAccess {
			canRead = true
			break
		}
	}

	if !canRead {
		return ErrDatabaseInfoDenied
	}

	return nil
}

func (s *Service) WriteInfo(ctx context.Context) error {
	return a10n.UserAuthorized(ctx, a10n.AdminWrite)
}

func (s *Service) ReadConnection(ctx context.Context, env string, dbID string) error {
	user, err := a10n.GetIdentity(ctx)
	if err != nil {
		return err
	}

	roles, rolesErr := s.repo.FindUserRolesByDatabaseID(ctx, user.OrgID, env, dbID)
	if rolesErr != nil {
		return fmt.Errorf("check database roles: %w", rolesErr)
	}
	if len(roles) == 0 {
		return ErrDatabaseAccessDenied
	}

	return nil
}

func (s *Service) ReadTable(ctx context.Context, env string, dbID string, tableName string, columns ...string) error {
	user, err := a10n.GetIdentity(ctx)
	if err != nil {
		return err
	}

	roles, err := s.repo.FindUserRolesByDatabaseID(ctx, user.OrgID, env, dbID)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrDatabaseAccessDenied, err.Error())
	}

	canRead := false
	for _, role := range roles {
		if role.CanReadTable(tableName, columns...) {
			canRead = true
			break
		}
	}

	if !canRead {
		return ErrTableAccessDenied
	}

	return nil
}

func (s *Service) Close() error {
	return s.repo.Close()
}
