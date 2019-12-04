BINDATA := pkg/server/bindata/templates.go
TMPL_FILES := $(shell find ./templates/ -type f | grep -v pkg/server/bindata/templates.go | sed 's/ /\\ /g')
GO_FILES := $(shell find ./ -type f | grep .go | sed 's/ /\\ /g')

default: clean build

build: dist/app

dist/app: $(GO_FILES) bindata
	go build -ldflags "-s -w" -o $@

bindata: pkg/server/bindata/templates.go

pkg/server/bindata/templates.go:
	go-bindata -o $@ -pkg bindata templates/...

clean:
	go clean -i ./... | true
	rm -f ./dist/*
	rm -f pkg/server/bindata/templates.go

deps:
	go get -u github.com/go-bindata/go-bindata/...
	go get -u github.com/codegangsta/gin
	go get -u golang.org/x/lint/golint

dev:
	go-bindata -dev -o pkg/server/bindata/templates.go -pkg bindata templates/...
	gin

install: deps
	go mod download

lint:
	golint ./...

.PHONY: clean build bindata lint install
