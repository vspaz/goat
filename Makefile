BINARY_NAME=goat

all: build
build:
	go build -o $(BINARY_NAME) examples/main.go

test:
	go test -race ./...

clean:
	rm -f $(BINARY_NAME)