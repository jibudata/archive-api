.DEFAULT: all
.PHONY: all


API_PROTO_FILES=$(shell find . -name \*.proto | grep -v third_party)

fmt: goimports
	@$(GOIMPORTS) -w -local github.com/jibudata $(shell go list -f {{.Dir}} ./...)

vet: ## Run go vet against code.
	go vet ./pkg/...

tidy:
	go mod tidy


.PHONY: init
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: api
api:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=. --go_opt=paths=source_relative \
           --go-errors_out=. --go-errors_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           --go-http_out=. --go-http_opt=paths=source_relative \
           --openapi_out=enum_type=string:. --openapi_opt=paths=source_relative \
           --validate_out=. --validate_opt=paths=source_relative --validate_opt=lang=go \
           --openapiv2_out . \
           --openapiv2_opt logtostderr=true \
           --openapiv2_opt json_names_for_fields=false \
           $(API_PROTO_FILES)


.PHONY: lint
lint: golangci-lint
	$(golangci-lint) run --timeout=5m

GOIMPORTS = $(shell pwd)/build/bin/goimports
.PHONY: goimports
goimports: ## Download kustomize locally if necessary.
	$(call go-get-tool,$(GOIMPORTS),golang.org/x/tools/cmd/goimports@v0.5.0)

golangci-lint = $(shell pwd)/build/bin/golangci-lint
.PHONY: golangci-lint
golangci-lint: ## Download golangci-lint locally if necessary.
	$(call go-get-tool,$(golangci-lint),github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2)

test:
	go test -v client/
	#go build -gcflags=all="-N -l" -o build/bin/archival-client client/main.go




PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/build/bin go install $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef
