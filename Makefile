.PHONY: build
build:
	go build -v -ldflags "-s -w" ./cmd/app

.PHONY: run
run:
	go run ./cmd/app

.DEFUALT_GOAL := build
