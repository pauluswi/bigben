SHELL=/bin/bash
GOMODULE=github.com/pauluswi/bigben

E2E_FLAG=-tags=e2e
INTEGRATION_FLAGS=-tags=integration

COVERAGE_OUT=coverage.cov
COVERAGE_FLAGS=-coverprofile=$(COVERAGE_OUT) -cover
COVERAGE_TOOL_FUNC_CMD=go tool cover -func $(COVERAGE_OUT)
COVERAGE_TOOL_HTML_CMD=go tool cover -html $(COVERAGE_OUT)

internal/auth/mocks/mock_repository.go:
	mockgen -destination=$@ -package=mocks $(GOMODULE)/internal/auth Repository

internal/auth/mocks/mock_usecase.go:
	mockgen -destination=$@ -package=mocks $(GOMODULE)/internal/auth Usecase

internal/healthcheck/mocks/mock_usecase.go:
	mockgen -destination=$@ -package=mocks $(GOMODULE)/internal/healthcheck Usecase

pkg/client/ocore/mocks/mock_client.go:
	mockgen -destination=$@ -package=mocks $(GOMODULE)/pkg/client/ocore Client

mockgen: internal/auth/mocks/mock_repository.go \
	internal/auth/mocks/mock_usecase.go \
	internal/healthcheck/mocks/mock_usecase.go \
	pkg/client/ocore/mocks/mock_client.go

test: export ENV := test
test:
	go test ./... -cover

test-integration: export ENV := test
test-integration: migrate-up
	go test $(INTEGRATION_FLAGS) ./...

test-e2e: export ENV := test
test-e2e: migrate-up
	go test $(E2E_FLAG) ./e2e_test/... -v -count=1

coverage: export ENV := test
coverage:
	go test $(COVERAGE_FLAGS) ./...
	$(COVERAGE_TOOL_FUNC_CMD)

coverage-html: export ENV := test
coverage-html:
	go test $(COVERAGE_FLAGS) ./...
	$(COVERAGE_TOOL_HTML_CMD)

coverage-integration: export ENV := test
coverage-integration: migrate-up
	go test $(INTEGRATION_FLAGS) $(COVERAGE_FLAGS) ./...
	$(COVERAGE_TOOL_FUNC_CMD)

coverage-integration-html: export ENV := test
coverage-integration-html: migrate-up
	go test $(INTEGRATION_FLAGS) $(COVERAGE_FLAGS) ./...
	$(COVERAGE_TOOL_HTML_CMD)

coverage-e2e: export ENV := test
coverage-e2e: migrate-up
	go test $(E2E_FLAG) $(COVERAGE_FLAGS) ./e2e_test/... -v -count=1
	$(COVERAGE_TOOL_FUNC_CMD)

coverage-e2e-html: export ENV := test
coverage-e2e-html: migrate-up
	go test $(E2E_FLAG) $(COVERAGE_FLAGS) ./e2e_test/... -v -count=1
	$(COVERAGE_TOOL_HTML_CMD)

run:
	go run .

clean:
	rm -v internal/auth/mocks/mock_*.go && \
	rm -v internal/healthcheck/mocks/mock_*.go

lint:
	golangci-lint run --print-issued-lines=false --exclude-use-default=false --enable=golint --enable=goimports  --enable=unconvert --enable=unparam --enable=gosec --timeout=10m

infra-up:
	docker-compose up -d

infra-down:
	docker-compose down

migration-script:
	@if [ "$(name)" = "" ]; then\
		printf '%s\n\n' '-- +migrate Up' '-- +migrate Down' > sql/migration/$(shell date '+%Y%m%d%H%M%S')_new_script.sql;\
	else\
		printf '%s\n\n' '-- +migrate Up' '-- +migrate Down' > sql/migration/$(shell date '+%Y%m%d%H%M%S')_$(name).sql;\
	fi

migrate-up:
	go run main.go migrate --direction=up --step=0

migrate-down:
	go run main.go migrate --direction=down --step=1

genkey:
	@go run main.go genkey $(ARGS)

pubkey:
	@go run main.go pubkey $(ARGS)

encrypt:
	@go run main.go encrypt $(ARGS)

decrypt:
	@go run main.go decrypt $(ARGS)

server:
	go run main.go server

docker-image:
	docker build \
		--build-arg BUILD_USER=$(BUILD_USER) \
		--build-arg BUILD_TOKEN=$(BUILD_TOKEN) \
		-t sabine/sabine .

