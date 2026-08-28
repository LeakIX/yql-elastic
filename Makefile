.PHONY: help
help: ## Ask for help!
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "%-20s %s\n", $$1, $$2}'

.PHONY: vet
vet: ## Run go vet
	go vet ./...

.PHONY: lint
lint: ## Run golangci-lint over the whole repo
	golangci-lint run

.PHONY: vuln
# Baked into gobuild:v9. Override to run it without installing it.
GOVULNCHECK ?= govulncheck

vuln: ## Scan for known vulnerabilities (vuln.go.dev)
	$(GOVULNCHECK) ./...

.PHONY: test
test: ## Run tests with benchmarks
	go test -v ./... -bench=.

.PHONY: check
check: vet test ## Run all checks (vet + test)

.PHONY: check-format
check-format: ## Check Go code formatting
	@unformatted=$$(gofmt -l .); \
	if [ -n "$$unformatted" ]; then \
		echo "The following files are not formatted:"; \
		echo "$$unformatted"; \
		exit 1; \
	fi

.PHONY: format
format: ## Format Go code
	gofmt -w .

.PHONY: check-changelog
check-changelog: ## Check every commit referenced in CHANGELOG.md exists
	./scripts/check-changelog-commits.sh