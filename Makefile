BINDATA := web/app/bindata/templates.go
TMPL_FILES := $(shell find ./web/templates/ -type f)
SLUG ?= $(shell git rev-parse --abbrev-ref HEAD)-$(shell git rev-parse HEAD|cut -c1-7)
WEB := $(shell find ./web/app -type f|grep .go |grep -v app/docs/bindata.go | grep -v bindata.go)

default: build

release: clean build

build: dist/app

dist/app:
	go generate ./...
	go build -ldflags "-s -w -X bgm38/config.Version=$(SLUG)" -o $@

clean:
	go clean -i ./... | true
	rm -f ./dist/*
	rm -f web/app/bindata/templates.go
	rm -rf web/app/docs/

install:
	go get github.com/swaggo/swag/cmd/swag
	go get github.com/go-bindata/go-bindata/...
	go mod download

.PHONY: clean build install
