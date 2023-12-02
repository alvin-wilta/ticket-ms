package main

import (
	"log"
	"net"

	"github.com/alvin-wilta/ticket-ms/ticket_service/db"
	"github.com/alvin-wilta/ticket-ms/ticket_service/nsq"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	nsqAddr := "localhost:4150"
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	db := db.InitDB()
	grpcServer := InitGRPC(db)
	handler := nsq.New(db)
	nsq.InitNSQConsumer("Create", nsqAddr, handler.HandleMessage)

	log.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
