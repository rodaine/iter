.PHONY: all
all: lint test

.PHONY: test
test:
	go test -race -cover

.PHONY: lint
lint:
	golangci-lint run