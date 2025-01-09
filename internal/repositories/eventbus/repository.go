package eventbus

import (
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
)

type repository struct {
	kafkaWriter ekafka.KafkaPubSub
}

func New(kafkaWriter ekafka.KafkaPubSub) *repository {
	return &repository{
		kafkaWriter: kafkaWriter,
	}
}
