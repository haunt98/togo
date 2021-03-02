.PHONY: help unittest lint run docker_up docker_down integration_test

help:
	@echo read Makefile

unittest:
	go test ./...

lint:
	golangci-lint run ./...

run:
	go run *.go

docker_up:
	docker-compose up -d

docker_down:
	docker-compose down

integration_test:
	go test -tags=integration ./...
