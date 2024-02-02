package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogapopp/microservices/chat-server/internal/handler"
)

func main() {
	grpcserver := handler.NewGRPCServer()
	listen, err := net.Listen("tcp", "localhost:8082")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err = grpcserver.Serve(listen)
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("chat server is running at :8082")
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigint
	grpcserver.GracefulStop()
}
