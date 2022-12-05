# Copyright 2022 clavinjune/ononoki-engine
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

include tools.mk

.PHONY: fmt
fmt:
	@gofmt -w -s .
	@goimports -w .
	@go mod tidy
	@go fix ./...
	@go run $(LICENSER) apply -r "clavinjune/ononoki-engine" 2> /dev/null

.PHONY: lint
lint:
	@go vet ./...
	@go run $(GOVULNCHECK) ./...
	@go run $(LICENSER) verify
	@go run $(LINTER) run

.PHONY: test
test:
	@go test -covermode=count -coverprofile out/cover.out -json ./... > out/test-report.json

.PHONY: test
test-cover: test
	@go tool cover -html=out/cover.out
