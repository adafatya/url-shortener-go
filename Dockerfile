# =========================
# Builder
# =========================
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY go.mod ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/api


# =========================
# Runtime
# =========================
FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/app ./app

EXPOSE 8080

CMD ["./app"]
