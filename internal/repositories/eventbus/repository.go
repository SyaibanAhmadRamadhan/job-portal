package eventbus

import (
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
)

type repository struct {
	client ekafka.KafkaPubSub
	conf   *conf.KafkaConfig
}

func New(client ekafka.KafkaPubSub, conf *conf.KafkaConfig) *repository {
	return &repository{
		client: client,
		conf:   conf,
	}
}
