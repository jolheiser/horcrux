GO ?= go
VERSION := $(shell git describe --tags --always | sed 's/-/+/' | sed 's/^v//')

.PHONY: lint
lint:
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		export BINARY="golangci-lint"; \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.24.0; \
	fi
	golangci-lint run --timeout 5m

.PHONY: vet
vet:
	$(GO) vet ./...

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: test
test:
	$(GO) test -race -v ./...

.PHONY: build
build: generate
	$(GO) build -ldflags '-s -w -X "go.jolheiser.com/horcrux/config.Version=$(VERSION)"'

.PHONY: generate
generate:
	$(GO) generate ./...