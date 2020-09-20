#!/usr/bin/make -f

test: fmt
	go test -count=1 ./...

fmt:
	go fmt ./...

build:
	go build ./...
