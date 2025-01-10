package eventbus

import (
	"context"
	"errors"
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"google.golang.org/protobuf/proto"
	"time"
)

func (r *repository) ConsumerJobPostETL(ctx context.Context) (output ekafka.SubOutput, err error) {
	mechanism := plain.Mechanism{
		Username: r.conf.Username,
		Password: r.conf.Password,
	}

	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}
	output, err = r.client.Subscribe(ctx, ekafka.SubInput{
		Unmarshal: func(bytes []byte, a any) error {
			payloadBuf, ok := a.(proto.Message)
			if !ok {
				return tracer.ErrInternalServer(errors.New("invalid payload, must be generated proto buf"))
			}
			return proto.Unmarshal(bytes, payloadBuf)
		},
		Config: kafka.ReaderConfig{
			Brokers: []string{r.conf.Host},
			GroupID: "ConsumerJobPostETLGroup1",
			Topic:   "jobpostetl",
			Dialer:  dialer,
		},
	})
	if err != nil {
		return output, tracer.ErrInternalServer(err)
	}

	return
}
