PKG_PREFIX=github.com/tbe-team/raybot/internal/build
VERSION=$(shell git describe --tags --abbrev=8 --dirty --always --long)
BUILD_DATE=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS:= \
	-X $(PKG_PREFIX).Version=$(VERSION) \
	-X $(PKG_PREFIX).Date=$(BUILD_DATE)

########################
# Code generation
########################
.PHONY: gen-openapi
gen-openapi:
	set -eux

	pnpm --package=@redocly/cli@1.34 dlx redocly bundle ./api/openapi/openapi.yml --output api/openapi/openapi.gen.yml --ext yml
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 \
		-config internal/handlers/http/gen/oapi-codegen.yml \
		api/openapi/openapi.gen.yml

.PHONY: gen-sqlc
gen-sqlc:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.28.0 generate --file internal/storage/db/sqlc/sqlc.yml

.PHONY: gen-mock
gen-mock:
	go run github.com/vektra/mockery/v2@v2.53.1 --config .mockery.yml

.PHONY: gen-all
gen-all: gen-openapi gen-sqlc gen-mock

#########################
# Database
#########################
GOOSE_DRIVER=sqlite3
GOOSE_DBSTRING="file:./bin/raybot.db"
GOOSE_MIGRATION_DIR=internal/storage/db/migration

.PHONY: migrate-up
migrate-up:
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) \
	go run github.com/pressly/goose/v3/cmd/goose@v3.24.1 up

.PHONY: migrate-down
migrate-down:
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) \
	go run github.com/pressly/goose/v3/cmd/goose@v3.24.1 down

.PHONY: migrate-status
migrate-status:
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) \
	go run github.com/pressly/goose/v3/cmd/goose@v3.24.1 status

.PHONY: migrate-create
migrate-create:
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) \
	go run github.com/pressly/goose/v3/cmd/goose@v3.24.1 create "$(name)" sql

.PHONY: migrate-reset
migrate-reset:
	GOOSE_DRIVER=$(GOOSE_DRIVER) \
	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
	GOOSE_MIGRATION_DIR=$(GOOSE_MIGRATION_DIR) \
	go run github.com/pressly/goose/v3/cmd/goose@v3.24.1 reset

#########################
# Build
#########################
.PHONY: build
build:
	CGO_ENABLED=1 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -ldflags "$(LDFLAGS)" -o bin/raybot cmd/raybot/main.go

.PHONY: build-ui
build-ui:
	make -C ui build

.PHONY: build-arm64
build-arm64:
	set -eux

	docker build \
		--build-arg PKG_PREFIX=$(PKG_PREFIX) \
		--build-arg VERSION=$(VERSION) \
		--build-arg BUILD_DATE=$(BUILD_DATE) \
		-t raybot-builder-deb11 \
		-f docker/raybot-build-deb11.dockerfile .
	docker create --name temp-build raybot-builder-deb11:latest
	docker cp temp-build:/app/raybot ./raybot-arm64
	docker rm temp-build

#########################
# Docker
#########################
.PHONY: docker-build-raybot
docker-build-raybot:
	docker build -t raybot -f docker/raybot.dockerfile .

.PHONY: docker-run-raybot
docker-run-raybot:
	docker run --rm -it \
		-v ./bin/config.yml:/app/config.yml \
		-v ./bin/raybot.db:/app/raybot.db \
		-p 3000:3000 \
		raybot \
		/app/raybot -config /app/config.yml -db /app/raybot.db

#########################
# Run
#########################
.PHONY: run
run:
	go run -ldflags "$(LDFLAGS)" cmd/raybot/main.go -config bin/config.yml -db bin/raybot.db

#########################
# Testing
#########################
.PHONY: test
test:
	go test -v --failfast ./...

.PHONY: test-cov
test-cov:
	go test -coverprofile=bin/coverage.out ./...
	go tool cover -html=bin/coverage.out -o bin/coverage.html
	@echo "Coverage report saved to bin/coverage.html"

########################
# Lint
########################
.PHONY: lint-go
lint-go:
	golangci-lint run ./... --config .golangci.yml
