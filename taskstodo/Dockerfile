# Usar uma imagem base do Go para desenvolvimento
FROM golang:1.22-alpine
WORKDIR /taskstodo 
# Copiar os arquivos de configuração do Go e baixar as dependências
COPY go.mod go.sum ./
RUN go mod download
# Copiar todos os arquivos do projeto
COPY . .
# Run the tests in the container

RUN go build -o serv-API ./cmd/
EXPOSE 8080
# Iniciar o servidor em de producao
CMD ["./serv-API"]