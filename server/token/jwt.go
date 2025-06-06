package token

import (
	"crypto/rsa"
	"fmt"
	"io"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	Values = map[string]any

	Settings struct {
		SigningMethod jwt.SigningMethod
		ExpTime       time.Duration
	}

	JWT struct {
		settings   Settings
		privateKey *rsa.PrivateKey
		publicKey  *rsa.PublicKey
	}

	claims struct {
		jwt.RegisteredClaims
		Attrs Values
	}

	Token struct {
		ExpiresAt time.Time
		Values    Values
	}
)

func New(privateReader, pubReader io.Reader, settings Settings) (*JWT, error) {
	privateBytes, err := io.ReadAll(privateReader)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	pubBytes, err := io.ReadAll(pubReader)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key: %w", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	return &JWT{
		privateKey: privateKey,
		publicKey:  pubKey,
		settings:   settings,
	}, nil
}

func (s *JWT) Create(attributes Values) (string, error) {
	t := jwt.New(s.settings.SigningMethod)
	now := time.Now()
	t.Claims = claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.settings.ExpTime)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Attrs: attributes,
	}

	return t.SignedString(s.privateKey)
}

func (s *JWT) Validate(tokenString string) (*Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	attrs, ok := claims["Attrs"].(Values)
	if !ok {
		return nil, fmt.Errorf("invalid token values")
	}

	expTime, err := claims.GetExpirationTime()
	if err != nil {
		return nil, fmt.Errorf("invalid exp time")
	}

	return &Token{
		ExpiresAt: expTime.Time,
		Values:    attrs,
	}, nil
}
