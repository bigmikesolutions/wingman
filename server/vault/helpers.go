package vault

import (
	"encoding/json"
	"fmt"
)

const (
	attrValue = "value"
)

type SecretValue map[string]any

func marshall(v any) (SecretValue, error) {
	j, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var s map[string]any
	if err := json.Unmarshal(j, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func unmarshall(secret SecretValue, v any) error {
	j, err := json.Marshal(secret)
	if err != nil {
		return fmt.Errorf("marshal back to JSON: %w", err)
	}
	if err := json.Unmarshal(j, &v); err != nil {
		return fmt.Errorf("unmarshal into target: %w", err)
	}
	return nil
}
