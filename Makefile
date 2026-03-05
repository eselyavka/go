GO_MODULE_DIR := localhost/leetcode
GOCACHE_DIR := /tmp/go-build
GOMODCACHE_DIR := $(CURDIR)/pkg/mod

.PHONY: fmt fmt-check test validate lint verify

fmt:
	cd $(GO_MODULE_DIR) && find . -type f -name '*.go' -print0 | xargs -0 gofmt -w

fmt-check:
	cd $(GO_MODULE_DIR) && test -z "$$(find . -type f -name '*.go' -print0 | xargs -0 gofmt -l)"

test:
	cd $(GO_MODULE_DIR) && GOCACHE=$(GOCACHE_DIR) GOMODCACHE=$(GOMODCACHE_DIR) go test ./...

validate:
	./scripts/validate_solutions.sh

lint:
	cd $(GO_MODULE_DIR) && GOCACHE=$(GOCACHE_DIR) GOMODCACHE=$(GOMODCACHE_DIR) golangci-lint run ./...

verify: fmt-check validate test
