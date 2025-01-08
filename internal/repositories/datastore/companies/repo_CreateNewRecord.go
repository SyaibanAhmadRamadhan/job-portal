package companies

import (
	"context"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
)

func (r *repository) CreateNewRecord(ctx context.Context, input CreateNewRecordInput) (output CreateNewRecordOutput, err error) {
	return
}

type CreateNewRecordInput struct {
	Tx   wsqlx.Rdbms
	Name string
}

type CreateNewRecordOutput struct {
	ID string
}
