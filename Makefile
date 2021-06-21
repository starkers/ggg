SHELL := /usr/bin/env bash
VERSION := $(shell git describe --always --dirty --tags 2> /dev/null || "undefined")
BINARY := ggg-${VERSION}
ECHO = echo -e

REPO ?= github.com/starkers/ggg

DEFAULT_GOAL:=help

## Print this help
#  eg: 'make' or 'make help'
help:
	@awk -v skip=1 \
		'/^##/ { sub(/^[#[:blank:]]*/, "", $$0); doc_h=$$0; doc=""; skip=0; next } \
		 skip  { next } \
		 /^#/  { doc=doc "\n" substr($$0, 2); next } \
		 /:/   { sub(/:.*/, "", $$0); printf "\033[1m%-30s\033[0m\033[1m%s\033[0m %s\n\n", $$0, doc_h, doc; skip=1 }' \
		$(MAKEFILE_LIST)

## Build the binary
#  (placed in root of repository)
build:
	env GOPRIVATE=$(REPO) CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-X main.VERSION=${VERSION} -s -w" -o $(BINARY)-Darwin-arm64 $(REPO)/
	env GOPRIVATE=$(REPO) CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.VERSION=${VERSION} -s -w" -o $(BINARY)-Darwin-amd64 $(REPO)/
	env GOPRIVATE=$(REPO) CGO_ENABLED=0 GOOS=linux  GOARCH=amd64 go build -ldflags="-X main.VERSION=${VERSION} -s -w" -o $(BINARY)-Linux-amd64  $(REPO)/

## Clean artifacts
#
clean:
	rm -v $(BINARY)-*

