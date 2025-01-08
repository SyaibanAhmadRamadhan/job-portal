package jobs

import "context"

type JobRepository interface {
	CreateNewRecord(ctx context.Context, input CreateNewRecordInput) (output CreateNewRecordOutput, err error)
}
