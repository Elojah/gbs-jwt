PACKAGE    = gbs-jwt
DATE      ?= $(shell date +%FT%T%z)
VERSION   ?= $(shell echo $(shell cat $(PWD)/.version)-$(shell git describe --tags --always))
DIR        = $(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

GO         = go
GOROOT     ?= $(shell go env GOROOT)
GODOC      = godoc
GOFMT      = gofmt

# For CI
ifneq ($(wildcard ./bin/golangci-lint),)
	GOLINT = ./bin/golangci-lint
else
	GOLINT = golangci-lint
endif

V          = 0
Q          = $(if $(filter 1,$V),,@)
M          = $(shell printf "\033[0;35m▶\033[0m")

GO_PACKAGE        = github.com/elojah/gbs-jwt
SECRETS              = secrets

GEN_PB_GO          = protoc -I=$(GOPATH)/src --gogoslick_out=$(GOPATH)/src
GEN_PB_SERVICE_GO  = protoc -I=$(GOPATH)/src --gogoslick_out=plugins=grpc,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:$(GOPATH)/src

.PHONY: all
all: secrets

.PHONY: secrets
secrets:  ## Build secrets binary
	$(info $(M) building executable secrets…) @
	$Q cd cmd/$(SECRETS) && $(GO) build \
		-mod=readonly \
		-tags release \
		-ldflags '-X main.version=$(VERSION)' \
		-o ../../bin/$(PACKAGE)_$(SECRETS)_$(VERSION)
	$Q cp bin/$(PACKAGE)_$(SECRETS)_$(VERSION) bin/$(PACKAGE)_$(SECRETS)

# Proto lang
.PHONY: proto-go
proto-go:    PB_LANG = GO
proto-go: ## Regenerate protobuf files
	$(info $(M) running protobuf $(PB_LANG)…) @
	$(info $(M) generate utils…) @
	$Q $(GEN_PB_$(PB_LANG)) github.com/gogo/protobuf/gogoproto/gogo.proto
	$(info $(M) generate domain…) @
	$Q $(GEN_PB_$(PB_LANG)) $(GO_PACKAGE)/pkg/jwt/jwt.proto
	$(info $(M) generate clients…) @
	$(info $(M) generate dto…) @
	$(info $(M) generate services…) @
	$Q $(GEN_PB_SERVICE_$(PB_LANG)) $(GO_PACKAGE)/cmd/$(SECRETS)/grpc/$(SECRETS).proto

# Proto
.PHONY: proto
proto: proto-go

# Vendor
.PHONY: vendor
vendor:
	$(info $(M) running go mod vendor…) @
	$Q $(GO) mod vendor

# Tidy
.PHONY: tidy
tidy:
	$(info $(M) running go mod tidy…) @
	$Q $(GO) mod tidy

# Check
.PHONY: check
check: vendor test lint

# Lint
.PHONY: lint
lint:
	$(info $(M) running $(GOLINT)…)
	$Q $(GOLINT) run

# Test
.PHONY: test
test:
	$(info $(M) running go test…) @
	$Q $(GO) test -cover -race -v ./...

# Clean
.PHONY: clean
clean:
	$(info $(M) cleaning bin…) @
	$Q rm -rf bin/$(PACKAGE)_$(SECRETS)_*

## Helpers

.PHONY: go-version
go-version: ## Print go version used in this makefile
	$Q echo $(GO)

.PHONY: fmt
fmt: ## Format code
	$(info $(M) running $(GOFMT)…) @
	$Q $(GOFMT) ./...

.PHONY: doc
doc: ## Generate project documentation
	$(info $(M) running $(GODOC)…) @
	$Q $(GODOC) ./...

.PHONY: version
version: ## Print current project version
	@echo $(VERSION)

.PHONY: help
help: ## Print this
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
