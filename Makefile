.PHONY:
.SILENT:

build:
    go build -o ./.bin/bot cmd/telega/main.go

run: build
    ./.bin/bot