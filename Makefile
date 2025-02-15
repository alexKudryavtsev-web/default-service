include .env
export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

compose-up-prod: ### PRODUCTION: run docker-compose
	docker-compose up --build -d postgres app

compose-down: ### UNIVERSAL: Down docker-compose
	docker-compose down --remove-orphans

compose-up-dev: ### DEVELOPMENT: Run postgres
	docker-compose up --build -d postgres && docker-compose logs -f

run-app: ### DEVELOPMENT: Run app (after `make compose-up-dev`)
	go run -tags migrate ./cmd/app

remove-volume: ### UNIVERSAL: remove docker volume
	docker volume rm defaultservice_pg-data

build: ### UNIVERSAL: build the application
	go build -tags migrate -o $(LOCAL_BIN)/app ./cmd/app