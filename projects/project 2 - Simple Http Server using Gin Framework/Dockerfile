# Stage 1 - build
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Install git (untuk download dependencies)
RUN apk add --no-cache git

# Copy go.mod dan go.sum dulu (agar caching download dependency)
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# ⬇️ INI PENTING: Copy file config
COPY config/config.json ./config/config.json

# Build binary
RUN go build -o myapp main.go

# Stage 2 - minimal image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .

COPY --from=builder /app/config ./config

EXPOSE 8080

CMD ["./myapp"]
