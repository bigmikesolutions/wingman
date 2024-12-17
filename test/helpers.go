package test

import (
	"context"
	"time"
)

const clientTimeout = 1 * time.Second

func testContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), clientTimeout)
}
