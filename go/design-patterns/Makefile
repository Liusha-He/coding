SHELL := /bin/bash

test:
	GODOG_FEATURES=./features go test -timeout 30s -v -cover ./tests/...

.PHONY: test
