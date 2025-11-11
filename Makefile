run:
	go run main.go

build:
	go build -o tictactoe main.go

install:
	go mod tidy

test: 
	go test ./...

all: run build test install

.PHONY: run build install test all