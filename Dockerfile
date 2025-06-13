FROM golang:1.24.3-alpine

# Install necessary packages
RUN apk add --no-cache \
    bash curl git make build-base unzip \
    && go install github.com/air-verse/air@latest \
    && go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application code
COPY . .

# Expose app port (adjust if needed)
EXPOSE 8080

# Enable direnv and run Air
CMD ["bash", "-c", "air"]
