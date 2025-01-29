FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o bittorrent

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bittorrent .

CMD ["./bittorrent"]