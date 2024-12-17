.DEFAULT_GOAL := default
.PHONY: default test

VER := $(shell git rev-parse --short HEAD)
BRANCH_TAG := "latest"
COVERAGE_FILE := "coverage.cov"
GO_OS := "linux"
GO_ARCH := "amd64"

default: build test

lint:
	@golangci-lint run ./...

fmt:
	@goimports -local "github.com/bigmikesolutions/wingman" -l -w .
	@gofumpt -l -w .

build: generate build-api

build-api:
	@CGO_ENABLED=0 go build -o bin/api \
		-ldflags "-X github.com/bigmikesolutions/wingman/pkg/build.Version=${VER}" \
		./cmd/api

install-gqlgen:
	@go get -d "github.com/99designs/gqlgen"

generate: vendor-delete install-gqlgen
	@go run github.com/99designs/gqlgen generate

vendor-delete:
	@rm -rf vendor

vendor:
	@go mod tidy
	@go mod vendor

test:
	@go test -short -race -count=1 -v ./...

local-run-api: build-api
	@./bin/api

local-docker: local-docker-stop local-docker-start

local-docker-start:
	@docker-compose up -d --build

local-docker-stop:
	@docker-compose stop

local-docker-down:
	@docker-compose down
