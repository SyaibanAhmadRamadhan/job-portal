package eventbus

import ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"

type repository struct {
	kafkaBroker ekafka.KafkaPubSub
}

func New(kafkaBroker ekafka.KafkaPubSub) *repository {
	return &repository{
		kafkaBroker: kafkaBroker,
	}
}
