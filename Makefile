.PHONY: build

build:
	go build -v ./cmd/api-server

.DEFAULT_GOAL := build
