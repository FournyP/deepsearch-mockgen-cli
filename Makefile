.PHONY: build test fmt mocks help

build:
	go build ./...

test:
	go test $$(go list ./tests/... | grep '_unit_tests')

fmt:
	gofmt -w .

mocks:
	go run . -S ./src -O ./tests -A -P

help:
	@echo "build  - go build ./..."
	@echo "test   - go test unit-test packages"
	@echo "fmt    - gofmt -w ."
	@echo "mocks  - regenerate mocks with this very CLI (needs mockgen on PATH and a TTY)"
