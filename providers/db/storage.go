package db

import (
	"context"
	"fmt"

	"github.com/bigmikesolutions/wingman/service/vault"
)

type (
	secureStorage interface {
		write(context.Context, ConnectionInfo) error
		read(context.Context, ID) (*ConnectionInfo, error)
	}

	VaultStorage struct {
		vault *vault.Secrets
	}
)

func NewStorage(vault *vault.Secrets) *VaultStorage {
	return &VaultStorage{
		vault: vault,
	}
}

func (v *VaultStorage) path(id ID) string {
	return fmt.Sprintf("/providers/db/connections/%s", id)
}

func (v *VaultStorage) write(ctx context.Context, info ConnectionInfo) error {
	s, err := vault.ToSecretValue(info)
	if err != nil {
		return err
	}

	return v.vault.Write(ctx, v.path(info.ID), s)
}

func (v *VaultStorage) read(ctx context.Context, id ID) (*ConnectionInfo, error) {
	s, err := v.vault.Read(ctx, v.path(id))
	if err != nil {
		return nil, err
	}

	info := &ConnectionInfo{}
	if err := vault.FromSecretValue(s, info); err != nil {
		return nil, err
	}
	return info, nil
}
