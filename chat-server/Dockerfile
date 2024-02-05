FROM golang:1.20.3-alpine AS builder

COPY . /github.com/gogapopp/microservices/chat-server/source/
WORKDIR /github.com/gogapopp/microservices/chat-server/source/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/chat-server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/gogapopp/microservices/chat-server/source/bin/chat_server .

CMD ["./chat_server"]