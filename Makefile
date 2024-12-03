APP=valkey-tpc

.PHONY: help all build clean

help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "    help   : Show help"
	@echo "    all    : Build multiple binary of this project"
	@echo "    build  : Build the binary of this project for current platform"
	@echo "    clean  : Clean the binary of this project"

all:build

build:
	@go build -o ${APP}

test:
	@go test

clean:
	@rm -rf ${APP}