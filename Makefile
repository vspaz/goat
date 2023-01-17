BINARY_NAME=goat

all: build
build:
	go build -o $(BINARY_NAME) examples/main.go

.PHONY: test
test:
	go test -race -v -p 4 ./...

.PHONY: clean
clean:
	rm -f $(BINARY_NAME)

.PHONY: style-fix
style-fix:
	gofmt -w .

.PHONY: lint
lint:
	golangci-lint run

.PHONY: upgrade
upgrade:
	go mod tidy
	go get -u all ./...
