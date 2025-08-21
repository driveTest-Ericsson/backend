# Stage 1: Build
FROM golang:1.24.3-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git build-base

WORKDIR /app

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go app (static binary)
RUN go build -o server .

# Stage 2: Run
FROM alpine:3.20

# Install CA certificates (needed for HTTPS calls)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server .

# Expose app port (adjust if needed)
EXPOSE 8080

# Run the binary
CMD ["./server"]
