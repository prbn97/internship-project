# Use a Go base image
FROM golang:1.22-alpine
WORKDIR /taskstodo 

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all files and build the binary
COPY . .
RUN go build -o serv-API ./cmd/
EXPOSE 8080
CMD ["./serv-API"]