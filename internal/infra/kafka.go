package infra

import (
	"context"
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func NewKafkaWriter(config conf.KafkaConfig) (ekafka.KafkaPubSub, primitive.CloseFunc) {
	mechanism := plain.Mechanism{
		Username: config.Username,
		Password: config.Password,
	}
	sharedTransport := &kafka.Transport{
		SASL: mechanism,
	}

	kafkaWriter := ekafka.New(ekafka.WithOtel(), ekafka.KafkaCustomWriter(&kafka.Writer{
		Addr:      kafka.TCP(config.Host),
		Balancer:  &kafka.LeastBytes{},
		Transport: sharedTransport,
	}))

	return kafkaWriter, func(ctx context.Context) error {
		kafkaWriter.Close()
		return nil
	}
}
