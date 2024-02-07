package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogapopp/microservices/chat-server/internal/handler"
	"github.com/gogapopp/microservices/chat-server/internal/repository/sqlite"
)

func main() {
	ctx := context.Background()
	db, err := sqlite.NewRepostiry(ctx, "./storage.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	grpcserver := handler.NewGRPCServer(db)
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err = grpcserver.Serve(listen)
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Printf("server listening at %v", listen.Addr())
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigint
	grpcserver.GracefulStop()
}
