package vault

import "encoding/json"

const (
	attrValue = "value"
)

func ToSecretValue(v any) (SecretValue, error) {
	j, err := json.Marshal(v)
	if err != nil {
		return SecretValue{}, err
	}

	s := make(SecretValue)
	s["value"] = string(j)
	return s, nil
}

func FromSecretValue(secret SecretValue, v any) error {
	b := []byte(secret[attrValue].(string))
	return json.Unmarshal(b, v)
}
