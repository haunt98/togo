.PHONY: help test lint

help:
	@echo make unittest
	@echo make lint
	@echo make run

unittest:
	go test ./...

integration_test:
	docker-compose up -d
	go test -tags=integration ./...

lint:
	golangci-lint run ./...

run:
	go run *.go
