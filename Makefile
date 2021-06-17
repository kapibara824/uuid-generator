.PHONY: all
all:
	make build
build:
	go build -o bin/idgen cmd/main.go