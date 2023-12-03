package main

import (
	"log"
	"net"

	"github.com/alvin-wilta/ticket-ms/ticket_service/config"
	"github.com/alvin-wilta/ticket-ms/ticket_service/db"
	"github.com/alvin-wilta/ticket-ms/ticket_service/nsq"
)

func main() {
	cfg := config.New()
	host := cfg.ServiceAddr + ":" + cfg.ServicePort
	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	db := db.InitDB(cfg)
	grpcServer := InitGRPC(db, cfg)
	nsqHandler := nsq.New(db, cfg)
	nsq.InitNSQConsumer("Create", cfg, nsqHandler.HandleMessage)

	log.Printf("Server is listening on port %v...", cfg.ServicePort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
