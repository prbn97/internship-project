# Usar uma imagem base do Node.js para desenvolvimento
# Definir diretorio
FROM node:22-alpine AS builder 
WORKDIR /app
# Copiar e instalar arquivos do npm e dependências
COPY package*.json ./
RUN npm install
# Copiar todos os arquivos do projeto
COPY . .
# Run the tests in the container
RUN npm run build
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
# Expor a porta usada pelo servidor de producao
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]