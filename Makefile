.PHONY: all requirements build run

ifndef BIN_NAME
    BIN_NAME = warlog
    ifndef GOOS
        ifeq ($(OS), Windows_NT)
            BIN_NAME := $(BIN_NAME).exe
        endif
    endif
    ifeq ($(GOOS), windows)
        BIN_NAME := $(BIN_NAME).exe
    endif
endif

all: build

requirements:
	@echo Installing requirements...
	@go get -v "github.com/gocraft/web"

build: requirements
	@echo Building...
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BIN_NAME) -v .

run: build
	@echo Running...
	@./$(BIN_NAME)