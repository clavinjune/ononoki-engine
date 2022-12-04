include tools.mk

.PHONY: fmt
fmt:
	@gofmt -w -s .
	@go mod tidy
	@go fix ./...

.PHONY: lint
lint:
	@go vet ./...
	@go run ${GOVULNCHECK} ./...

.PHONY: test
test:
	@go test -covermode=count -coverprofile out/cover.out -json ./... > out/test-report.json

.PHONY: test
test-cover: test
	@go tool cover -html=out/cover.out