package nsqw

import "github.com/nsqio/go-nsq"

type HandlerNSQ struct {
	producer *nsq.Producer
}

func New(producer *nsq.Producer) *HandlerNSQ {
	return &HandlerNSQ{
		producer: producer,
	}
}
