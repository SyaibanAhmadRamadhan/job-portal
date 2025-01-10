package eventbus

import (
	"context"
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
)

type PublisherRepository interface {
	PublishJobPostETL(ctx context.Context, input PublishJobPostETLInput) (err error)
}

type ConsumerRepository interface {
	ConsumerJobPostETL(ctx context.Context) (output ekafka.SubOutput, err error)
}
