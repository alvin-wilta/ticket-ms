package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alvin-wilta/ticket-ms/proxy_service/graph"
)

const defaultPort = "50001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// NOTE: Initialize NSQ
	nsqHandler := initNSQHandler()
	// NOTE: Initialize gRPC
	client := initGRPC()

	// NOTE: Initialize GQL handler
	resolver := graph.InitResolver(*client, *nsqHandler)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		log.Print("[HTTP] Healthcheck ok")
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ ", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
