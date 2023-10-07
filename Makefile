.PHONY: clean all init  generate_mocks test, migrate, start, start-non-binary

# Variables
DB_URI := mongodb://localhost:27017
DB_NAME := diPayDB

all: init migrate build


build: cmd/main.go  generated
	@echo "Building..."
	go build -o $@ $<

init:
	go mod tidy
	go mod vendor

migrate:
	go run cmd/main.go migrate

test:
	go test -short -coverprofile coverage.out -v ./...

coverage: test
	go tool cover -func=coverage.out

start-binary:
	./build

start-non-binary:
	go run cmd/main.go

generated: cmd/dipay.yml
	@echo "Generating files..."
	oapi-codegen -package=api -generate "types,server,spec" cmd/dipay.yml > api/api.gen.go


generate_mocks: $(INTERFACES_GEN_GO_FILES)
$(INTERFACES_GEN_GO_FILES): %.mock.gen.go: %.go
	@echo "Generating mocks $@ for $<"
	mockgen -source=$< -destination=$@ -package=$(shell basename $(dir $<))