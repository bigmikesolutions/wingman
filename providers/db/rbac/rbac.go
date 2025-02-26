package rbac

import (
	"context"
	"fmt"
)

type (
	userRoleRepo interface {
		CreateUserRole(ctx context.Context, role UserRole) error
		FindUserRolesByDatabaseID(ctx context.Context, id string) ([]UserRole, error)
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
	// TODO check user rights
	return s.repo.CreateUserRole(ctx, role)
}

func (s *Service) ReadInfo(ctx context.Context, dbID string) error {
	roles, err := s.repo.FindUserRolesByDatabaseID(ctx, dbID)
	if err != nil {
		return ErrDatabaseAccessDenied
	}

	canRead := false
	for _, role := range roles {
		if role.Info != nil && *role.Info == ReadOnlyAccess {
			canRead = true
			break
		}
	}

	if !canRead {
		return ErrTableAccessDenied
	}

	return nil
}

func (s *Service) WriteInfo(ctx context.Context) error {
	// TODO check user rights
	return nil
}

func (s *Service) ReadConnection(ctx context.Context, dbID string) error {
	roles, rolesErr := s.repo.FindUserRolesByDatabaseID(ctx, dbID)
	if rolesErr != nil {
		return fmt.Errorf("check database roles: %w", rolesErr)
	}
	if len(roles) == 0 {
		return ErrDatabaseAccessDenied
	}

	return nil
}

func (s *Service) ReadTable(ctx context.Context, dbID string, tableName string, columns ...string) error {
	roles, err := s.repo.FindUserRolesByDatabaseID(ctx, dbID)
	if err != nil {
		return ErrDatabaseAccessDenied
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
