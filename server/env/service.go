package env

import (
	"context"
	"errors"

	"github.com/bigmikesolutions/wingman/server/a10n"
)

type (
	repo interface {
		FindByID(ctx context.Context, id ID) (*Environment, error)
		Create(ctx context.Context, env Environment) error
	}

	Service struct {
		repo repo
	}
)

var ErrNotFound = errors.New("environment not found")

func New(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindByID(ctx context.Context, id ID) (*Environment, error) {
	if err := a10n.UserAuthenticated(ctx); err != nil {
		return nil, err
	}

	return s.repo.FindByID(ctx, id)
}

func (s *Service) Create(ctx context.Context, env Environment) error {
	if err := a10n.UserAuthorized(ctx, a10n.AdminWrite); err != nil {
		return err
	}

	return s.repo.Create(ctx, env)
}
