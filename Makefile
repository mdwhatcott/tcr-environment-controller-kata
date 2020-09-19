#!/usr/bin/make -f

test: fmt
	go test -count=1 -v ./...

fmt:
	go fmt ./...

build:
	go build ./...
