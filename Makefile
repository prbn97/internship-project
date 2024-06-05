BINARY_NAME=serv-API

.PHONY: run
run: build
	./$(BINARY_NAME)

.PHONY: build
build:
	go build -o $(BINARY_NAME) ./cmd/

.PHONY: test
test:
	go test ./... -race -cover