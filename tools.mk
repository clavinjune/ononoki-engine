GOIMPORTS_VERSION := v0.4.0
GOIMPORTS := golang.org/x/tools/cmd/goimports@$(GOIMPORTS_VERSION)

GOVULNCHECK_VERSION := latest
GOVULNCHECK := golang.org/x/vuln/cmd/govulncheck@$(GOVULNCHECK_VERSION)

LICENSER_VERSION = v0.6.0
LICENSER := github.com/liamawhite/licenser@$(LICENSER_VERSION)

LINTER_VERSION = v1.50.1
LINTER   := github.com/golangci/golangci-lint/cmd/golangci-lint@$(LINTER_VERSION)