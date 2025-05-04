FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server_fio ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o /fio_service ./cmd/server

FROM debian:bookworm-slim
RUN set -x && apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /fio_service /app/fio_service
COPY ./config ./config
COPY ./.env .
COPY ./migrations ./migrations
COPY ./docs ./docs

EXPOSE 8080

CMD ["./fio_service"]