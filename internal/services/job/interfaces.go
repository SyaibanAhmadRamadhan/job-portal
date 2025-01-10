package job

import "context"

type JobService interface {
	GetListJob(ctx context.Context, input GetListJobInput) (output GetListJobOutput, err error)
	CreateJob(ctx context.Context, input CreateJobInput) (output CreateJobOutput, err error)
	ConsumerPostJobETL(ctx context.Context) (err error)
}
