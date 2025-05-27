# Dockerfile
FROM --platform=linux/amd64 golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o order-packs-calculator ./cmd/server

FROM --platform=linux/amd64 alpine:3.16

WORKDIR /app

COPY --from=builder /app/order-packs-calculator /app/
COPY --from=builder /app/web /app/web

EXPOSE 8080

CMD ["/app/order-packs-calculator"]