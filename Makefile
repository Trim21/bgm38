GENERATED  := client.exe
BINDATA := pkg/server/bindata/tmpl_bindata.go pkg/server/bindata/gzipped_static_bindata.go
STATIC_FILES := $(shell find ./bindata/ -type f | grep gzipped | grep min | sed 's/ /\\ /g')
TMPL_FILES := $(shell find ./bindata/ -type f | grep -v gzipped | sed 's/ /\\ /g')
GO_FILES := $(shell find ./ -type f | grep .go | sed 's/ /\\ /g')
BINDATA_OPTS := '-ignore="\\.DS_Store" -pkg pkg/bindata'

build: dist/app

dist/app: $(GO_FILES)
	go build -ldflags "-s -w" --tags jsoniter -o $@

bindata: $(BINDATA)

pkg/server/bindata/tmpl_bindata.go: $(TMPL_FILES)
	go-bindata -o $@ $(BINDATA_OPTS) $(TMPL_FILES)

pkg/server/bindata/gzipped_static_bindata.go: $(STATIC_FILES)
	bindata -o $@ $(BINDATA_OPTS) $(STATIC_FILES)

clean:
	go clean -i ./... | true
	rm -f ./dist/*
	rm -f ./pkg/bindata/*

deps:
	GO111MODULE=off go get -u github.com/shuLhan/go-bindata/cmd/go-bindata
	GO111MODULE=off go install github.com/shuLhan/go-bindata/cmd/go-bindata
	GO111MODULE=off go get -u github.com/kataras/bindata/cmd/bindata
	GO111MODULE=off go install github.com/kataras/bindata/cmd/bindata
	GO111MODULE=off go get -u github.com/codegangsta/gin
	GO111MODULE=off go install github.com/codegangsta/gin
	GO111MODULE=off go get -u golang.org/x/lint/golint
	GO111MODULE=off go install golang.org/x/lint/golint

dev:
	DEV=1 gowatch

lint:
	golint ./...
