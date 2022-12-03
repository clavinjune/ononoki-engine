.PHONY: fmt
fmt:
	@gofmt -w -s .
	@go fix ./...

.PHONY: lint
lint:
	@go vet ./...
	@go run golang.org/x/vuln/cmd/govulncheck@latest ./...

.PHONY: test
test:
	@go test -covermode=count -coverprofile out/cover.out -json ./... > out/test-report.json

.PHONY: test
test-cover: test
	@go tool cover -html=out/cover.out