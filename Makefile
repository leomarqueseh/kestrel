.PHONY: run build test fmt vet

run:
	go run ./cmd/api

build:
	go build -o bin/kestrel ./cmd/api

test:
	go test ./... -v

fmt:
	go fmt ./...

vet:
	go vet ./...
