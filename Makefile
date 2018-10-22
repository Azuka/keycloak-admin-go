################################################################
## Color definition
################################################################
NO_COLOR    = \x1b[0m
OK_COLOR    = \x1b[32;01m
WARN_COLOR  = \x1b[33;01m
ERROR_COLOR = \x1b[31;01m

################################################################
## Helpers
################################################################
PWD := $(shell pwd)
GO_PACKAGES = $(shell go list ./... | grep -v vendor | grep -v integration)
GO_INTEGRATION_PACKAGES = $(shell go list ./... | grep integration)
GO_FILES = $(shell find . -name "*.go" | grep -v vendor | uniq)
CI_TEST_REPORTS ?= /tmp/test-results


.PHONY: init-ci
init-ci:
	@echo "$(OK_COLOR)==> Installing minimal build requirements$(NO_COLOR)"
	dep ensure -v
	curl -L https://git.io/vp6lP | sh
	go get -u github.com/jstemmer/go-junit-report

.PHONY: init
init: init-ci
	@echo "$(OK_COLOR)==> Installing dev build requirements$(NO_COLOR)"
	go get -u github.com/fatih/gomodifytags
	go get -u github.com/mailru/easyjson/...

# Format files
.PHONY: format
format:
	@echo "$(OK_COLOR)==> Formatting$(NO_COLOR)"
	gofmt -s -l -w $(GO_FILES)

# Generate files
.PHONY: generate
generate:
	@echo "$(OK_COLOR)==> Generating code$(NO_COLOR)"
	@rm -rf $(PWD)/keycloak/*_easyjson.go
	@go generate ./...

# Lint
.PHONY: lint
lint:
	@echo "$(OK_COLOR)==> Linting$(NO_COLOR)$(NO_COLOR)"
	@go list -f '{{.Dir}}' ./... | grep -v 'vendor' | \
	xargs gometalinter

# Test
.PHONY: test
test: format lint
	@echo "$(OK_COLOR)==> Testing $(NO_COLOR)"
	go test -race -cover $(GO_PACKAGES)

# Test integration
.PHONY: integration
integration:
	@echo "$(OK_COLOR)==> Integration testing $(NO_COLOR)"
	docker-compose -f $(PWD)/integration/docker-compose.yml up -d
	go test -race $(GO_INTEGRATION_PACKAGES)
	make integration-clean

.PHONY: integration-clean
integration-clean:
	docker-compose -f $(PWD)/integration/docker-compose.yml down

# Generate coverage
.PHONY: coverage
coverage:
	@echo "$(OK_COLOR)==> Generating Coverage Report$(NO_COLOR)"
	mkdir -p $(CI_ARTIFACTS)/htmlcov
	overalls -project=$(PKG) -covermode=count
	go tool cover -html=overalls.coverprofile -o $(COVER_HTML)

# CI test
.PHONY: test-ci
#code coverage regex for Gitlab: ^total:(\s+)\(statements\)(\s+)(\d+(?:\.\d+)?%)
test-ci:
	@echo "$(OK_COLOR)==> Running ci test$(NO_COLOR)"
	mkdir -p $(CI_TEST_REPORTS)
	/bin/bash -c "set -euxo pipefail; \
	    go test -v -short -race -cover -coverprofile .testCoverage.txt $(GO_PACKAGES) | tee >(go-junit-report > $(CI_TEST_REPORTS)/report.xml); \
	    sed '/_easyjson.go/d' .testCoverage.txt > .testCoverage.txt.bak; mv .testCoverage.txt.bak .testCoverage.txt; go tool cover -func=.testCoverage.txt"

# CircleCI test
.PHONY: test-circle
test-circle:
	@echo "$(OK_COLOR)==> Running circle test$(NO_COLOR)"
	mkdir -p $(CI_TEST_REPORTS)
	/bin/bash -c "set -euxo pipefail; \
	    go test -v -short -race -cover $(GO_PACKAGES) | tee >(go-junit-report > $(CI_TEST_REPORTS)/report.xml)"

# CI Lint
.PHONY: lint-ci
lint-ci:
	@echo "$(OK_COLOR)==> Running CI lint$(NO_COLOR)"
	go list -f '{{.Dir}}' ./... | grep -v 'vendor' | xargs gometalinter --json > lint.json

# Quick test for rapid dev-feedback cycles
.PHONY: qt
qt:
	@echo "$(OK_COLOR)==> Running quick test$(NO_COLOR)"
	go test -short $(GO_PACKAGES)

.PHONY: local-ci
local-ci:
	@echo "$(OK_COLOR)==> Running CI locally. Did you run brew install gitlab-runner and CircleCI?$(NO_COLOR)"
	circleci local execute --job go-1.9
	circleci local execute --job go-1.10
	circleci local execute --job go-1.11
	brew services start gitlab-runner
	gitlab-runner exec docker unit
	gitlab-runner exec docker lint
