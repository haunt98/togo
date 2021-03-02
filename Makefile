.PHONY: help test lint

help:
	@echo make test: testing
	@echo make lint: linting

test:
	go test ./...

lint:
	golangci-lint run ./...
