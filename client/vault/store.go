package vault

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/zalando/go-keyring"
)

const (
	service = "wingman"
)

var ErrNotFound = errors.New("not found")

type Store struct{}

func New() *Store {
	return &Store{}
}

func (s Store) SetValue(key string, token any) error {
	b, err := json.Marshal(&token)
	if err != nil {
		return err
	}

	return keyring.Set(service, key, string(b))
}

func (s Store) GetValue(key string, token any) error {
	v, err := keyring.Get(service, key)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return ErrNotFound
		}

		return err
	}

	return json.Unmarshal([]byte(v), &token)
}
