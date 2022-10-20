.PHONY: install test-dev test cover run-dev

install:
	go mod download

test-dev:
	mkdir -p ./coverage && \
		go test -v -coverprofile=./coverage/coverage.out -covermode=atomic ./...

cover: test-dev
	go tool cover -func=./coverage/coverage.out &&\
		go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html

run-dev:
	go run ./cmd/cli/main.go

build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/quiz_master ./cmd/cli 