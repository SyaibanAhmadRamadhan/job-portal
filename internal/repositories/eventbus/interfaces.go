package eventbus

import "context"

type PublisherRepository interface {
	PublishJobPostETL(ctx context.Context, input PublishJobPostETLInput) (err error)
}
