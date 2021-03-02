.PHONY: help test lint

help:
	@echo make test: testing
	@echo make lint: linting
	@echo make run: running

test:
	go test ./...

lint:
	golangci-lint run ./...

run:
	go run *.go
