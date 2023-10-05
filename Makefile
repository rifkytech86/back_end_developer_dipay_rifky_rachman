.PHONY: clean all init generate generate_mocks

all: build/main

build/main: cmd/main.go generated
	@echo "Building..."
	go build -o $@ $<


init: generate
	go mod tidy
	go mod vendor


generated: cmd/dipay.yml
	@echo "Generating files..."
	oapi-codegen -package=api -generate "types,server,spec" cmd/dipay.yml > api/api.gen.go

