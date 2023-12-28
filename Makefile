.PHONY:
.SILENT:

build-go:
	go build -o ./.bin/bot cmd/main.go

run-go: build-go
	./.bin/bot

build:
	docker-compose build

up: build
	docker-compose up -d

down:
	docker-compose down

down-v:
	docker-compose down -v