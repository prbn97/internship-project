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

.PHONY: migration
migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@go run cmd/migrate/main.go up
	
.PHONY: migrate-down
migrate-down:
	@go run cmd/migrate/main.go down