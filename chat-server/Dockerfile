FROM golang:1.21-alpine AS builder

RUN apk add --no-cache --update go gcc g++

COPY . /github.com/gogapopp/microservices/chat-server/source/
WORKDIR /github.com/gogapopp/microservices/chat-server/source/

RUN go mod download
RUN CGO_ENABLED=1 go build -o ./bin/chat_server cmd/chat-server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/gogapopp/microservices/chat-server/source/bin/chat_server .

CMD ["./chat_server"]