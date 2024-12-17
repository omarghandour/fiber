# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder

# Environment setup
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY ./api /app

# Build the Go binary
RUN go build -o main .

# Stage 2: Minimal runtime image
FROM alpine:latest

# Working directory for runtime
WORKDIR /root/

# Copy the built binary
COPY --from=builder /app/main .

# Expose the port your Fiber app listens to
EXPOSE 3000

# Command to start the app
CMD ["./main"]
