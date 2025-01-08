package index_jobs

import "context"

type JobRepository interface {
	GetAll(ctx context.Context, input GetAllInput) (output GetAllOutput, err error)
}
