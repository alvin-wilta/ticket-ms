package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/alvin-wilta/ticket-ms/proxy_service/config"
	"github.com/alvin-wilta/ticket-ms/proxy_service/nsqw"
	"github.com/nsqio/go-nsq"
)

func initNSQHandler(cfg *config.Config) *nsqw.HandlerNSQ {
	config := nsq.NewConfig()
	addr := fmt.Sprintf("%s:%s", cfg.NsqAddr, cfg.NsqPort)
	producer, err := nsq.NewProducer(addr, config)
	if err != nil {
		log.Fatalf("[NSQ] Initialization error: %v", err)
	}

	// Gracefully handle interrupt signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		producer.Stop()
		os.Exit(0)
	}()
	handler := nsqw.New(producer)
	return handler
}
