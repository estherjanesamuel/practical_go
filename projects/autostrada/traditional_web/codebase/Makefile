# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format code and tidy modfile
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go test -race -vet=off ./...
	go mod verify


# ==================================================================================== #
# BUILD
# ==================================================================================== #

## build: build the cmd/web application
.PHONY: build
build:
	go mod verify
	go build -ldflags='-s' -o=./bin/web ./cmd/web
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o=./bin/linux_amd64/web ./cmd/web

## run: run the cmd/web application
.PHONY: run
run:
	go run github.com/cosmtrek/air@v1.40.4 --c="./air.toml"


# ==================================================================================== #
# SQL MIGRATIONS
# ==================================================================================== #

## migrations/new name=$1: create a new database migration
.PHONY: migrations/new
migrations/new:
	go run -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest create -seq -ext=.sql -dir=./assets/migrations ${name}

## migrations/up: apply all up database migrations
.PHONY: migrations/up
migrations/up:
	go run -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="sqlite3://db.sqlite" up

## migrations/down: apply all down database migrations
.PHONY: migrations/down
migrations/down:
	go run -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="sqlite3://db.sqlite" down

## migrations/goto version=$1: migrate to a specific version number
.PHONY: migrations/goto
migrations/goto:
	go run -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="sqlite3://db.sqlite" goto ${version}

## migrations/force version=$1: force database migration
.PHONY: migrations/force
migrations/force:
	go run -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="sqlite3://db.sqlite" force ${version}

## migrations/version: print the current in-use migration version
.PHONY: migrations/version
migrations/version:
	go run -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path=./assets/migrations -database="sqlite3://db.sqlite" version


