SHELL := /bin/bash

test:
	go test ./tests/

run:
	CGO_ENABLED=0 go build -o design ./src/cmd
	./design

.PHONY: test run
