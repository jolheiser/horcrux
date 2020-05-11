GO ?= go

.PHONY: lint
lint:
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		export BINARY="golangci-lint"; \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin v1.23.1; \
	fi
	golangci-lint run --timeout 5m

.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: test
test:
	$(GO) test -race -v ./...

.PHONY: build
build: generate
	$(GO) build

.PHONY: generate
generate:
	$(GO) generate ./...