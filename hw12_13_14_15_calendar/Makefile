ifneq (,$(wildcard ./.env))
    include .env
    export
endif

BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" -tags=jsoniter ./cmd/calendar

run: build
	$(BIN) -config ./configs/config.toml

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

build-app:
	docker-compose up --build -d -e LDFLAGS="$(LDFLAGS)"

down-app:
	docker-compose down

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

migrate:
	migrate -path ./migrations -database 'postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):5432/$(DB_DATABASE)?sslmode=disable' up

format:
	gofumpt -l -w .

test:
	go test -race ./internal/...

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.37.0

lint: install-lint-deps
	golangci-lint run ./...

.PHONY: build run build-img run-img version test lint
