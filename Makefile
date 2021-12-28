SHELL=/bin/bash
GOMODULE=github.com/pauluswi/bigben

infra-up:
	docker-compose up --build

infra-down:
	docker-compose down

migrate-up:
	go run main.go migrate up

migrate-down:
	go run main.go migrate down

seed:
	go run main.go seed CustomerSeed AccountSeed

unit-test:
	go test -v controller/*.go -race -coverprofile=coverage.out -covermode=atomic

serve:
	go run main.go serve

docker-image:
	docker build \
		--build-arg BUILD_USER=$(BUILD_USER) \
		--build-arg BUILD_TOKEN=$(BUILD_TOKEN) \
		-t bigben/bigben .