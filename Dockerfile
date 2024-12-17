# Use official Go image for building the project
FROM golang:1.20-alpine AS builder

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code (api folder content)
COPY ./api /app

# Build the binary
RUN go build -o main .

# Minimal final image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose port 3000 for Fiber
EXPOSE 3000

# Command to run the app
CMD ["./main"]
