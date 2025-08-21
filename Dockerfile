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
RUN go build -buildvcs=false -o ./bin/main ./cmd/api

# Stage 2: Run
FROM alpine:3.20

# Install CA certificates (needed for HTTPS calls)
RUN apk add --no-cache bash ca-certificates curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz \
    | tar xvz -C /usr/local/bin \
    && chmod +x /usr/local/bin/migrate

WORKDIR /app

# Copy only the binary from builder
COPY --from=builder /app/bin/main .
COPY --from=builder /app/Makefile ./Makefile
COPY ./entrypoint.sh ./entrypoint.sh
COPY ./cmd/migrate/migrations ./cmd/migrate/migrations

# Expose app port (adjust if needed)
EXPOSE 8080

# Run the binary
CMD ["sh", "./entrypoint.sh"]
