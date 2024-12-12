# Backend (Go) Build Stage
FROM golang:1.23.4 AS backend-build

# Set working directory for backend
WORKDIR /backend

# Copy backend code and dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./

# Build the Go backend
RUN go build -o main .

# Frontend (React) Build Stage
FROM node:18 AS frontend-build

# Set working directory for frontend
WORKDIR /frontend

# Add NODE_OPTIONS to fix OpenSSL issue
ENV NODE_OPTIONS=--openssl-legacy-provider

# Copy frontend dependencies and install
COPY frontend/package*.json ./
RUN npm install

# Copy the rest of the frontend code
COPY frontend/ ./

# Build the React app
RUN npm run build

# Final Stage: Combine Backend and Frontend
FROM alpine:latest

# Install lightweight web server (e.g., Caddy for serving static files)
RUN apk add --no-cache caddy

# Set working directory in the final image
WORKDIR /app

# Copy backend binary and environment variables
COPY --from=backend-build /backend/main /app/backend
COPY backend/.env /app/

# Copy frontend build files
COPY --from=frontend-build /frontend/build /app/frontend/build

# Expose backend and frontend ports
EXPOSE 8080
EXPOSE 3000

# Command to run the backend binary
CMD ["/app/backend"]
