.PHONY: dev wire

VERSION = 0.1.0
BIN_PATH = ./bin/marine-tracer

build:
	go build -ldflags "-w -s -X main.VERSION=$(VERSION)" -o $(BIN_PATH) ./cmd

dev:
	go run -ldflags "-X main.VERSION=$(VERSION)" ./cmd/main.go s -c ./configs/config.yml -m ./configs/casbin_model.conf

watch:
	air -c .air.conf

wire:
	wire gen ./internal/app/core/injector/wire.go
