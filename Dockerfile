FROM golang:1.25-bookworm AS builder
WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main .

FROM debian:bookworm-slim
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
