# Start the production phase
## Use a Node.js base image
FROM node:22-alpine AS builder 
WORKDIR /app
## Copy and install npm files and dependencies
COPY package*.json ./
RUN npm install
## Copy all files build
COPY . .
RUN npm run build
# Start the production server
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]