# Stage 1: Build the Go binary
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o rate-limiter ./cmd/server

# Stage 2: Minimal image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/rate-limiter .

EXPOSE 8080

CMD ["./rate-limiter"]
