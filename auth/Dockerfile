FROM golang:1.21 AS builder

COPY . /github.com/gogapopp/microservices/auth/source/
WORKDIR /github.com/gogapopp/microservices/auth/source/

RUN go mod download
RUN CGO_ENABLED=1 go build -o ./bin/auth_server cmd/auth/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/gogapopp/microservices/auth/source/bin/auth_server .

CMD ["./auth_server"]