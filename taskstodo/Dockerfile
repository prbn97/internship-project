# Use a Go base image
FROM golang:1.22-alpine
WORKDIR /taskstodo 

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download
# Add necessary build tools
RUN apk add --no-cache gcc musl-dev

# Copy all files and build the binary
COPY . .
ENV CGO_ENABLED=1
RUN go build -o serv-API ./cmd/
EXPOSE 8080
CMD ["./serv-API"]