SLUG ?= $(shell git rev-parse --abbrev-ref HEAD)-$(shell git rev-parse HEAD|cut -c1-7)
MSGP_GEN = pkg/log/model_gen.go
DOC = pkg/web/docs/swagger.json pkg/web/docs/swagger.yaml pkg/web/docs/docs.go
SRC = $(filter-out $(DOC) $(MSGP_GEN), $(shell find . -type f -name '*.go'))
WEB_SRC = $(filter-out $(DOC) $(MSGP_GEN), $(shell find ./pkg/web/ -type f -name '*.go'))
ASSERTS = $(wildcard asserts/**/* asserts/*)

default: build

release: clean build

build: dist/app

doc: $(DOC)
generated: $(MSGP_GEN) $(DOC) pkg/asserts/pkged.go

$(MSGP_GEN): %_gen.go: %.go
	msgp -file $<

$(DOC): $(WEB_SRC)
	swag init --generalInfo ./pkg/web/doc.go -o ./pkg/web/docs

pkg/asserts/pkged.go: $(ASSERTS)
	rm -f $@
	pkger -include /asserts -o pkg/asserts

dist/app: generated
	go build -mod=readonly -ldflags "-s -w -X bgm38/config.Version=$(SLUG)" -o $@

clean:
	go clean -i ./... | true
	rm -f ./dist/*
	rm -f $(DOC)
	rm -f $(MSGP_GEN)

deps:
	go mod download
	go get github.com/tinylib/msgp
	go get github.com/swaggo/swag/cmd/swag
	go get github.com/markbates/pkger/cmd/pkger

install: deps generated

.PHONY: clean build deps generated install
