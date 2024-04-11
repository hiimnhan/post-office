BUILD_DIR := tmp/bin/post-office

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## tidy: format and tidy go files
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## build: build the project
.PHONY: build
build:
	go build -o ${BUILD_DIR}

## run: run the project
.PHONY: run
run: build
	${BUILD_DIR}

## run/debug: run the project in debug mode
.PHONY: run/debug
run/debug:
	go run main.go --debug
