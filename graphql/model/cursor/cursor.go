package cursor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

type (
	Cursor string

	Value struct {
		Offset int `json:"offset"`
	}
)

func Encode(v Value) (Cursor, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", fmt.Errorf("cursor encode: %w", err)
	}

	b64 := base64.StdEncoding.EncodeToString(b)
	return Cursor(b64), nil
}

func Decode(cursor *Cursor) (Value, error) {
	if cursor == nil {
	}

	c := string(*cursor)
	if c == "" {
		return Value{}, nil
	}

	b, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return Value{}, fmt.Errorf("cursor decode: %w", err)
	}

	var v Value
	err = json.Unmarshal(b, &v)
	if err != nil {
		return Value{}, fmt.Errorf("cursor decode: %w", err)
	}

	return v, nil
}

func (c *Cursor) UnmarshalGQL(v any) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("cursor must be string")
	}
	*c = Cursor(value)
	return nil
}

func (c Cursor) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(fmt.Sprintf(`"%s"`, c)))
}
