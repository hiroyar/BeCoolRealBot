.PHONY:
.SILENT:

build-go:
	go build -o ./.bin/bot cmd/main.go

run: build-go
	./.bin/bot

