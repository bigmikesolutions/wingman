SHELL = /usr/bin/env bash
BINDIR 	?= $(CURDIR)/bin

BINNAME			= service
TESTFLAGS 	= -v -race -v
BUILDFLAGS	= -v -mod=vendor -trimpath
LDFLAGS   	= -w -s

ifdef CI_COMMIT_SHA
	GIT_COMMIT = ${CI_COMMIT_SHA}
endif
GIT_COMMIT ?= $(shell git rev-parse HEAD)

ifdef CI_COMMIT_TAG
	VERSION ?= ${CI_COMMIT_TAG}
endif
VERSION		?= $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)

ifeq ($(VERSION),)
	VERSION = $(shell git rev-parse --short HEAD)
endif

GIT_DIRTY		:= $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")
APP_MODULE	:= $(shell go mod edit -print | head -1 | cut -d' ' -f2)

ifneq ($(VERSION),)
	LDFLAGS += -X ${APP_MODULE}/internal/version.version=${VERSION}
endif
LDFLAGS += -X ${APP_MODULE}/internal/version.timestamp=$(shell date +%s)
LDFLAGS += -X ${APP_MODULE}/internal/version.gitCommit=${GIT_COMMIT}
LDFLAGS += -X ${APP_MODULE}/internal/version.gitTreeState=${GIT_DIRTY}
ifneq ($(EXTRA_LDFLAGS),)
	LDFLAGS += ${EXTRA_LDFLAGS}
endif

.PHONY: build test vendor

all: build test

go.mod go.sum:
	go mod tidy

vendor: go.mod go.sum
	go mod vendor

vendor := $(shell test -d "vendor" && echo "" || echo "vendor")
build-client: $(vendor)
	CGO_ENABLED=0 go build $(BUILDFLAGS) -ldflags '$(LDFLAGS)' -o '$(BINDIR)/$(BINNAME)' ./cmd/client
