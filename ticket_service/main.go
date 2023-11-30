package main

import (
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	db := InitDB()
	grpcServer := InitGRPC(db)

	log.Println("Server is listening on port 50051...")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
