# Project settings
BINARY_NAME=mars_rover
SRC=cmd/main.go

.PHONY: all build run test tidy clean deps lint coverage

## Build the binary
build:
	go build -o $(BINARY_NAME) $(SRC)

## Run the program (interactively)
run:
	go run $(SRC)

## Run all unit tests
test:
	go test -v ./...

## Install dependencies (like testify)
deps:
	go get github.com/stretchr/testify

## Go tidy to clean up go.mod/go.sum
tidy:
	go mod tidy

## Clean built binary
clean:
	rm -f $(BINARY_NAME)

## Test with coverage (optional)
coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out
