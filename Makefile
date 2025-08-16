ENV_VARS =$(shell grep -v "^\#" .env | xargs)

envvars:
	@echo "export $(ENV_VARS)"

build:
	go build -o .exe main.go

.PHONY: server
server: envvars build
	./.exe server

.PHONY: mocks
mocks:
	mockgen -source=./internal/ports/repository.go -destination=./internal/ports/repository_mock.go -package=ports