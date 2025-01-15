package test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/bigmikesolutions/wingman/test/containers"
)

const (
	defaultTimeout     = 5 * time.Second
	dockerSetupTimeout = 120 * time.Second
)

var dc *containers.Service

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), dockerSetupTimeout)
	defer cancel()

	var err error
	dc, err = containers.New(ctx)
	if err != nil {
		panic(err)
	}
	defer dc.Close()

	code := m.Run()
	defer os.Exit(code)
}

type logger struct{}

func (c logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
