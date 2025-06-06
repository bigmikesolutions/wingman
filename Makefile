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

build: build-server build-cli

build-server:
	@CGO_ENABLED=0 go build -o bin/server \
		-ldflags "-X github.com/bigmikesolutions/wingman/service/build.Version=${VER}" \
		./cmd/server

build-cli:
	@CGO_ENABLED=0 go build -o bin/cli \
		./cmd/cli

generate:
	@go generate ./...

vendor-update:
	@go get -u ./...
	@go mod tidy
	@go mod vendor

vendor:
	@go mod tidy
	@go mod vendor

test:
	@go test -short -race -count=1 -v ./...

local-docker: local-terraform-clean-up local-docker-down local-docker-up

local-docker-up:
	@docker compose up -d --build

local-docker-down:
	@docker compose down

local-terraform-clean-up:
	@rm -rf terraform/.terraform* 2>/dev/null || true
	@rm terraform/terraform.tfstate 2>/dev/null || true

local-docker-wingman:
	@docker compose build
	@docker compose up -d  --force-recreate wingman
