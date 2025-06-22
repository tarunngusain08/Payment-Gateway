# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o payment-gateway ./cmd

# Run stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/payment-gateway .
COPY internal/config/config.yaml ./internal/config/config.yaml

EXPOSE 8080

CMD ["./payment-gateway"]
