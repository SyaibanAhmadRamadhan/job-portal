package job

import "context"

func (s *service) CreateJob(ctx context.Context, input CreateJobInput) (output CreateJobOutput, err error) {
	return
}

type CreateJobInput struct {
	CompanyName string
	Title       string
	Description string
}

type CreateJobOutput struct {
	CompanyID string
	JobID     string
}
