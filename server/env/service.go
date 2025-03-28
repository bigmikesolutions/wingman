package env

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/bigmikesolutions/wingman/server/a10n"
)

type (
	repo interface {
		FindByID(ctx context.Context, orgID OrganisationID, id ID) (*Environment, error)
		Create(ctx context.Context, env Environment) error
	}

	Service struct {
		repo repo
	}
)

var (
	ErrAlreadyExists = errors.New("environment already exists")
	ErrInvalid       = errors.New("invalid")
)

func New(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindByID(ctx context.Context, id ID) (*Environment, error) {
	user, err := a10n.GetIdentity(ctx)
	if err != nil {
		return nil, err
	}

	if err := user.ContainsRole(a10n.AdminRead, a10n.DeveloperRead); err != nil {
		return nil, err
	}

	if err := ValidateSession(ctx); err != nil {
		return nil, err
	}

	return s.repo.FindByID(ctx, user.OrgID, id)
}

func (s *Service) Create(ctx context.Context, e Environment) error {
	user, err := a10n.GetIdentity(ctx)
	if err != nil {
		return err
	}

	if err := user.ContainsRole(a10n.AdminWrite); err != nil {
		return err
	}

	e.OrgID = user.OrgID
	e.CreatedAt = time.Now().UTC()
	e.CreatedBy = user.UserID

	if err := e.Validate(); err != nil {
		return fmt.Errorf("%w: %s", ErrInvalid, err)
	}

	return s.repo.Create(ctx, e)
}
