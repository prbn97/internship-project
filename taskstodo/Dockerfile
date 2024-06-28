# Start the production server
## Use a Go base image
FROM golang:1.22-alpine
WORKDIR /taskstodo 
## Copy and install npm files and dependencies
COPY go.mod go.sum ./
RUN go mod download
RUN apk add --no-cache gcc musl-dev
## Copy all files build
COPY . .
## enable cgo
ENV CGO_ENABLED=1
## build
RUN go build -o serv-API ./cmd/
EXPOSE 8080
CMD ["./serv-API"]