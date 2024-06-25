FROM golang:1.22-alpine

WORKDIR /internship-project 
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o serv-API ./cmd/
EXPOSE 8080
CMD ["./serv-API"]