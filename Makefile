SLUG ?= $(shell git rev-parse --abbrev-ref HEAD)-$(shell git rev-parse HEAD|cut -c1-7)
MSGP_GEN = pkg/log/model_gen.go

default: build

release: clean build

build: dist/app

generated: $(MSGP_GEN) $(WEB)

$(MSGP_GEN): %_gen.go: %.go
	msgp -file $<

dist/app: generated
	go build -ldflags "-s -w -X bgm38/config.Version=$(SLUG)" -o $@

$(WEB): $(WEB_SRC)
	swag init --generalInfo ./web/app/doc.go -o ./web/app/docs

clean:
	go clean -i ./... | true
	rm -f ./dist/*

deps:
	go mod download
	go get github.com/tinylib/msgp

install: deps generated

.PHONY: clean build deps generated install
