FROM golang:1.27.1 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o calculator-api ./cmd/api

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/calculator-api .

EXPOSE 8080

CMD ["./calculator-api"]