package eventbus

import (
	"context"
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/proto/job_post_etl_payload"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"google.golang.org/protobuf/proto"
	"time"
)

func (r *repository) PublishJobPostETL(ctx context.Context, input PublishJobPostETLInput) (err error) {
	payloadByte, err := proto.Marshal(input.Payload)
	if err != nil {
		return tracer.ErrInternalServer(err)
	}

	randomKey := uuid.New().String()
	msg := kafka.Message{
		Topic:   "jobpostetl",
		Key:     []byte(randomKey),
		Value:   payloadByte,
		Headers: make([]kafka.Header, 0),
		Time:    time.Now().UTC(),
	}

	propagator := otel.GetTextMapPropagator()
	propagator.Inject(ctx, ekafka.NewMsgCarrier(&msg))

	_, err = r.client.Publish(ctx, ekafka.PubInput{
		Messages: []kafka.Message{msg},
	})
	if err != nil {
		return tracer.ErrInternalServer(err)
	}
	return
}

type PublishJobPostETLInput struct {
	Payload *job_post_etl_payload.JobPost
}
