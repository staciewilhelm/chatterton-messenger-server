#!make
include .env

clean-code:
	go fmt && go vet

deps:
	go mod

start-clean: reset-db clean-code
	go run main.go

start: clean-code
	go run main.go

# Database Migrations
reset-db: setup-migrations
	migrate -database ${POSTGRESQL_URL} -path migrations down
	migrate -database ${POSTGRESQL_URL} -path migrations up

setup-migrations:
	export POSTGRESQL_URL='postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_DATABASE)?sslmode=disable'

migrate-down: setup-migrations
	migrate -database ${POSTGRESQL_URL} -path migrations down

migrate-up: setup-migrations
	migrate -database ${POSTGRESQL_URL} -path migrations up

# Tests
run-tests:
	go test -ginkgo.randomizeAllSpecs

setup-tests:
	export PATH=$PATH:$(go env GOPATH)/bin
	ginkgo bootstrap
