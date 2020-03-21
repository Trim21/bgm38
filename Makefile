SLUG ?= $(shell git rev-parse --abbrev-ref HEAD)-$(shell git rev-parse HEAD|cut -c1-7)

default: build

release: clean build

build: dist/app

dist/app:
	go generate ./...
	go build -ldflags "-s -w -X bgm38/config.Version=$(SLUG)" -o $@

clean:
	go clean -i ./... | true
	rm -f ./dist/*

install:
	go mod download

.PHONY: clean build install
