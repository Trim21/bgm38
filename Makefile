COMMAND := go build -mod=readonly -ldflags "-s -w -X bgm38/config.Version=$(SLUG)" -o dist/app

COMMIT_SHA ?= $(shell git rev-parse HEAD)
COMMIT_SHORT_SHA ?= $(shell echo $(COMMIT_SHA)|cut -c1-7)
COMMIT_DATE ?= $(shell git show -s --pretty=%cs $(COMMIT_SHA))
COMMIT_REF ?= $(shell git rev-parse --abbrev-ref HEAD)
SLUG ?= $(COMMIT_REF)-$(COMMIT_SHORT_SHA)
MSGP_GEN = pkg/log/model_gen.go
DOC = pkg/web/docs/swagger.json pkg/web/docs/swagger.yaml pkg/web/docs/docs.go
SRC = $(filter-out $(DOC) $(MSGP_GEN), $(shell find -type f -name "*.go"))
WEB_SRC = $(filter-out $(DOC) $(MSGP_GEN), $(shell find pkg/web/ -type f -name "*.go"))
ASSERTS = $(wildcard asserts/**/* asserts/*)

default: build
	@echo 'git ref:    ' "'$(COMMIT_REF)'"
	@echo 'commit date:' "'$(COMMIT_DATE)'"
	@echo 'git hash:   ' "'$(COMMIT_SHORT_SHA)'"

release: clean generated
	$(COMMAND) -tags=asserts

build: dist/app

dist/app: generated
	@$(COMMAND)

generated: $(MSGP_GEN) $(DOC) pkg/asserts/pkged.go

$(MSGP_GEN): %_gen.go: %.go
	msgp -file $< -tests=false

$(DOC): $(WEB_SRC)
	@swag init --generalInfo ./pkg/web/doc.go -o ./pkg/web/docs --generatedTime=false

pkg/asserts/pkged.go: $(ASSERTS)
	@rm -f $@
	@pkger -include /asserts -o pkg/asserts

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
