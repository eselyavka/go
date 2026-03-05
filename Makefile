GO_MODULE_DIR := .
GOCACHE_DIR := /tmp/go-build
GOMODCACHE_DIR := $(CURDIR)/pkg/mod
GOPATH_DIR := /tmp/go

.PHONY: fmt fmt-check test validate lint verify

fmt:
	cd $(GO_MODULE_DIR) && find problems util -type f -name '*.go' -print0 | xargs -0 gofmt -w

fmt-check:
	cd $(GO_MODULE_DIR) && test -z "$$(find problems util -type f -name '*.go' -print0 | xargs -0 gofmt -l)"

test:
	cd $(GO_MODULE_DIR) && GO111MODULE=on GOPATH=$(GOPATH_DIR) GOCACHE=$(GOCACHE_DIR) GOMODCACHE=$(GOMODCACHE_DIR) go test ./problems/... ./util/...

validate:
	./scripts/validate_solutions.sh

lint:
	cd $(GO_MODULE_DIR) && GO111MODULE=on GOPATH=$(GOPATH_DIR) GOCACHE=$(GOCACHE_DIR) GOMODCACHE=$(GOMODCACHE_DIR) golangci-lint run ./problems/... ./util/...

verify: fmt-check validate test
