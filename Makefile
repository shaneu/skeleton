shell := /bin/bash

build:
	go build -o skeleton main.go

test:
	go test -race -v ./...