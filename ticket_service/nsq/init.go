package nsq

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/alvin-wilta/ticket-ms/ticket_service/config"
	"github.com/nsqio/go-nsq"
)

func InitNSQConsumer(channel string, cfg *config.Config, handler nsq.HandlerFunc) {
	nsqAddr := fmt.Sprintf("%v:%v", cfg.NsqAddr, cfg.NsqPort)

	config := nsq.NewConfig()
	config.MaxAttempts = uint16(cfg.NsqMaxAttempts)
	config.MaxInFlight = cfg.NsqMaxInFlight

	consumer, err := nsq.NewConsumer("ticketCreate", channel, config)
	if err != nil {
		log.Fatalf("[NSQ] Init consumer for topic ticket channel %v failed: %v", channel, err)
	}
	consumer.AddHandler(NewHandler(handler))

	err = consumer.ConnectToNSQD(nsqAddr)
	if err != nil {
		log.Fatalf("[NSQ] Failed connecting to NSQD on %v for channel %v", nsqAddr, channel)
	}

	// NOTE: Gracefully stop the consumer.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	go func() {
		<-signalChan
		consumer.Stop()
		os.Exit(0)
	}()

}
