# Usar uma imagem base do Go para desenvolvimento
FROM golang:1.22-alpine
# Instalar as dependências necessárias para cgo
RUN apk add --no-cache gcc musl-dev
WORKDIR /taskstodo 

COPY go.mod go.sum ./
RUN go mod download
# Copiar todos os arquivos do projeto
COPY . .
# Variável de ambiente para habilitar cgo
ENV CGO_ENABLED=1

RUN go build -o serv-API ./cmd/
EXPOSE 8080
# Iniciar o servidor em de producao
CMD ["./serv-API"]