package nsq

import (
	"log"

	"github.com/nsqio/go-nsq"
)

func InitNSQConsumer(channel string, nsqAddr string, handler nsq.HandlerFunc) {
	config := nsq.NewConfig()
	config.MaxAttempts = 3
	config.MaxInFlight = 2

	consumer, err := nsq.NewConsumer("ticket", channel, config)
	if err != nil {
		log.Fatalf("[NSQ] Init consumer for topic ticket channel %v failed: %v", channel, err)
	}
	consumer.AddHandler(NewHandler(handler))

	err = consumer.ConnectToNSQLookupd(nsqAddr)
	if err != nil {
		log.Fatalf("[NSQ] Failed connecting to NSQD on %v for channel %v", nsqAddr, channel)
	}

}
