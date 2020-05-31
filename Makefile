COMMIT_SHA ?= $(shell git rev-parse HEAD)
COMMIT_SHORT_SHA ?= $(shell echo $(COMMIT_SHA)|cut -c1-7)
COMMIT_DATE ?= $(shell git show -s --pretty=%cs $(COMMIT_SHA))
COMMIT_REF ?= $(shell git rev-parse --abbrev-ref HEAD)
SLUG ?= $(COMMIT_REF)-$(COMMIT_SHORT_SHA)
DOC = app/web/docs/swagger.json app/web/docs/swagger.yaml app/web/docs/docs.go
SRC = $(filter-out $(DOC), $(shell find -type f -name "*.go"))
WEB_SRC = $(filter-out $(DOC), $(shell find app/web/ -type f -name "*.go"))
ASSERTS = $(wildcard asserts/**/* asserts/*)
COMMAND := go build -mod=readonly -ldflags "-s -w -X bgm38/config.SHA=$(COMMIT_SHORT_SHA) -X bgm38/config.Version=$(SLUG)" -o dist/app

default: build
	@echo 'git ref:    ' "'$(COMMIT_REF)'"
	@echo 'commit date:' "'$(COMMIT_DATE)'"
	@echo 'git hash:   ' "'$(COMMIT_SHORT_SHA)'"

release: clean generated
	$(COMMAND) -tags=asserts

build: dist/app

dist/app: generated
	@$(COMMAND)

generated: $(DOC) pkg/asserts/pkged.go

$(DOC): $(WEB_SRC)
	@swag init --generalInfo ./app/web/doc.go --output ./app/web/docs --generatedTime=false

pkg/asserts/pkged.go: $(ASSERTS)
	@rm -f $@
	pkger -include /asserts -o pkg/asserts

clean:
	go clean -i ./... | true
	rm -f ./dist/*
	rm -f $(DOC)
	rm -f $(MSGP_GEN)

deps:
	go mod download
	go get github.com/swaggo/swag/cmd/swag
	go get github.com/markbates/pkger/cmd/pkger

install: deps generated

.PHONY: clean build deps generated install
