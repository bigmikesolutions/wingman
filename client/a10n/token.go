package a10n

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GetExpirationDate(t string) (*time.Time, error) {
	tokenClaims := jwt.MapClaims{}
	_, _, err := jwt.NewParser().ParseUnverified(t, &tokenClaims)
	if err != nil {
		return nil, err
	}

	expTime, err := tokenClaims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	return &expTime.Time, nil
}
