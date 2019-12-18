BINDATA := web/app/bindata/templates.go
TMPL_FILES := $(shell find ./web/templates/ -type f | sed 's/ /\\ /g')
GO_FILES := $(shell find ./ -type f | grep .go | sed 's/ /\\ /g')
GIT_VERSION := $(shell git describe --abbrev=0 --tags)-$(shell git rev-parse HEAD|cut -c1-8)

default: clean build

release: clean build

build: dist/app

dist/app: $(GO_FILES) bindata
	go build -ldflags "-s -w -X bgm38/config.Version=$(GIT_VERSION)" -o $@

bindata: web/app/bindata/templates.go

web/app/bindata/templates.go:
	go-bindata -o $@ -pkg bindata web/templates/...

clean:
	go clean -i ./... | true
	rm -f ./dist/*
	rm -f web/app/bindata/templates.go

dev:
	go-bindata -dev -o web/app/bindata/templates.go -pkg bindata web/templates/...
	gowatch

install:
	go get github.com/go-bindata/go-bindata/...
	go get github.com/silenceper/gowatch
	go get golang.org/x/lint/golint
	go mod download

lint:
	golint ./...

.PHONY: clean build bindata lint install
