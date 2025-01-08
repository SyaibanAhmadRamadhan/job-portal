package eventbus

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/proto/job_post_etl_payload"
)

func (r *repository) PublishJobPostETL(ctx context.Context, input PublishJobPostETLInput) (err error) {
	return
}

type PublishJobPostETLInput struct {
	Payload *job_post_etl_payload.JobPost
}
