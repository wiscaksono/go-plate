.PHONY: setup run build

setup:
	go mod download

run:
	go run cmd/main.go

build:
	go build -o bin/main cmd/main.go
