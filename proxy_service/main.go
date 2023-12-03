package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alvin-wilta/ticket-ms/proxy_service/config"
	"github.com/alvin-wilta/ticket-ms/proxy_service/graph"
)

func main() {
	cfg := config.New()

	// NOTE: Initialize NSQ
	nsqHandler := initNSQHandler(cfg)

	// NOTE: Initialize gRPC
	client := initGRPC(cfg)

	// NOTE: Initialize GQL handler
	resolver := graph.InitResolver(*client, *nsqHandler)
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		log.Print("[HTTP] Healthcheck ok")
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://%s:%s/ ", cfg.ServiceAddr, cfg.ServicePort)
	log.Fatal(http.ListenAndServe(cfg.ServiceAddr+":"+cfg.ServicePort, nil))
}
