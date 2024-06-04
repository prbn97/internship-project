BINARY_NAME=api-serv

.PHONY: run
run: build
	./$(BINARY_NAME)

.PHONY: build
build:
	go build -o $(BINARY_NAME) ./api/

.PHONY: test
test:
	go test ./... -race -cover