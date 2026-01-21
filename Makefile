APP_NAME = ecommerce-api
IMAGE = $(APP_NAME):latest

.PHONY: build run test lint docker-build docker-run

build:
	go build -v -o bin/$(APP_NAME) ./cmd/api

run:
	go run ./cmd/api

dev:
	air -c .air.toml

test:
	go test ./... -v

lint:
	gofmt -w .
	go vet ./...

docker-build:
	docker build -t $(IMAGE) .

docker-run:
	docker-compose up --build
docker-stop:
	docker-compose down