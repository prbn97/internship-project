# Use a Node.js base image
FROM node:22-alpine AS builder 
WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm install
#  Copy all files and build
COPY . .
RUN npm run build

# Use an Nginx base image for serving the build
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]