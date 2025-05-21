# Build stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Install necessary tools for building Go app
RUN apk add --no-cache git bash curl build-base

# Copy go module files and env
COPY go.mod go.sum ./
COPY .env .env

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go binary
RUN go build -o app .

# Final stage - run the app
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Install necessary tools (like bash, if needed)
RUN apk add --no-cache bash

# Copy the Go binary and .env file from the builder stage
COPY --from=builder /app/app /app/
COPY --from=builder /app/.env /app/

# Expose the necessary port
EXPOSE 5000

# Run the Go app directly
CMD ["./app"]
