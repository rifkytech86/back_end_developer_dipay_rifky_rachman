.PHONY: clean all init generate generate_mocks test

all: build/main

build/main: cmd/main.go generated
	@echo "Building..."
	go build -o $@ $<


init: generate
	go mod tidy
	go mod vendor

test:
	go test -short -coverprofile coverage.out -v ./...

generated: cmd/dipay.yml
	@echo "Generating files..."
	oapi-codegen -package=api -generate "types,server,spec" cmd/dipay.yml > api/api.gen.go


generate_mocks: $(INTERFACES_GEN_GO_FILES)
$(INTERFACES_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))