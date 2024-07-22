## Use a Go base image
FROM docker-hub.artifactory.globoi.com/golang:1.22-alpine
WORKDIR /taskstodo 
## Copy and install npm files and dependencies
COPY go.mod go.sum ./
RUN go mod download

## Copy all files build
COPY . .

## build for tsuru
ENV GOOS=linux GOARCH=amd64 
RUN go build -o serv_api ./cmd/
EXPOSE 8080
CMD ["./serv_api"]