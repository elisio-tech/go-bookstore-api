FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /app/bin/bookstore ./cmd/api

FROM alpine:3.21

RUN addgroup -S appgroup && adduser -S -G appgroup appuser \
    && apk add --no-cache tzdata ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/bookstore /app/bookstore

RUN mkdir -p /app/data && chown -R appuser:appgroup /app

USER appuser

EXPOSE 3000

CMD ["/app/bookstore"]