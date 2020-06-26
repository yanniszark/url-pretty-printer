.EXPORT_ALL_VARIABLES:
DOCKER_BUILDKIT		:= 1
GO111MODULE			:= on
CGO_ENABLED         ?= 0 
GOOS                ?= linux
GOARCH              ?= amd64

all: build

.PHONY=build
build:
	go build -a -o bin/url-pprint .