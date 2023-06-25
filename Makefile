.PHONY: test run build clean wire

APP_NAME = apiserver
GO ?= GO111MODULE=on go
BUILD_DIR = $(PWD)/build
MAIN_FILE = .//cmd/main.go
SERVER_FILE = ./internal/server.go
MIGRATION_DIR = $(PWD)/migrations

setup:
	go mod tidy
	sudo go install github.com/google/wire/cmd/wire@latest
	sudo go install github.com/swaggo/swag/cmd/swag@latest

test:
	go test ./...

# remove binary		
clean:
	echo "remove bin exe"
	rm -rf $(BUILD_DIR)

# build binary
build:
	CGO_ENABLED=0 $(GO) build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_FILE)

swag:
	swag init -g $(SERVER_FILE)

wire:
	cd internal && wire

# local run
local:
	make swag
	make wire
	make build
	$(BUILD_DIR)/$(APP_NAME)
