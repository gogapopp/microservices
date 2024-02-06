package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogapopp/microservices/auth/internal/handler"
	"github.com/jackc/pgx/v4"
)

const (
	dbDSN = "host=localhost port=5432 dbname=authdb user=authuser password=authpass sslmode=disable"
)

func main() {
	ctx := context.Background()
	grpcserver := handler.NewGRPCServer()
	listen, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err = grpcserver.Serve(listen)
		if err != nil {
			log.Fatal(err)
		}
	}()

	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer con.Close(ctx)

	log.Printf("server listening at %v", listen.Addr())
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigint
	grpcserver.GracefulStop()
}
