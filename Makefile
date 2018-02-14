.PHONY: build

ifndef BIN_NAME
BIN_NAME=warlog
endif

all: build

requirements:
	@go get "github.com/gocraft/web"

build: requirements
	@go build -o $(BIN_NAME) .
	@chmod +x ./$(BIN_NAME)