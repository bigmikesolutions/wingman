package vault

import (
	"encoding/json"
)

const (
	attrValue = "value"
)

type SecretValue map[string]any

func marshall(v any) (SecretValue, error) {
	j, err := json.Marshal(v)
	if err != nil {
		return SecretValue{}, err
	}

	s := make(SecretValue)
	s[attrValue] = string(j)
	return s, nil
}

func unmarshall(secret SecretValue, v any) error {
	b := []byte(secret[attrValue].(string))
	return json.Unmarshal(b, &v)
}
