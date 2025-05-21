# Single-stage build and run
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Install necessary tools for building and running Go app
RUN apk add --no-cache git bash curl build-base

# Copy go module files and env
COPY go.mod go.sum .env ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o app .

# Expose the application port
EXPOSE 5000

# Run the built binary
CMD ["./app"]
