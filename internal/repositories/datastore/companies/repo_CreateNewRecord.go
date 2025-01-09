package companies

import (
	"context"
	"errors"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/google/uuid"
)

func (r *repository) CreateNewRecord(ctx context.Context, input CreateNewRecordInput) (output CreateNewRecordOutput, err error) {
	if input.Tx == nil {
		return output, tracer.ErrInternalServer(errors.New("database transaction is not open"))
	}
	id := uuid.NewString()

	query := r.sq.Insert("companies").Columns(
		"id", "name",
	).Values(
		id, input.Name,
	).Suffix(
		`ON CONFLICT(name) DO UPDATE
				SET name=EXCLUDED.name
			 RETURNING id`,
	)

	err = input.Tx.QueryRowSq(ctx, query, wsqlx.QueryRowScanTypeStruct, &output)
	if err != nil {
		return output, tracer.ErrInternalServer(err)
	}
	return
}

type CreateNewRecordInput struct {
	Tx   wsqlx.Rdbms
	Name string
}

type CreateNewRecordOutput struct {
	ID string
}
