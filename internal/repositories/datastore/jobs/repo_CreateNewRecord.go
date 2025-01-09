package jobs

import (
	"context"
	"errors"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/google/uuid"
	"time"
)

func (r *repository) CreateNewRecord(ctx context.Context, input CreateNewRecordInput) (output CreateNewRecordOutput, err error) {
	if input.Tx == nil {
		return output, tracer.ErrInternalServer(errors.New("database transaction is not open"))
	}

	id := uuid.NewString()

	query := r.sq.Insert("jobs").Columns(
		"id", "company_id", "title", "description",
	).Values(
		id, input.CompanyID, input.Title, input.Description,
	).Suffix("RETURNING timestamp")

	err = input.Tx.QueryRowSq(ctx, query, wsqlx.QueryRowScanTypeStruct, &output)
	if err != nil {
		return output, tracer.ErrInternalServer(err)
	}

	output.ID = id
	return
}

type CreateNewRecordInput struct {
	Tx          wsqlx.Rdbms
	CompanyID   string
	Title       string
	Description string
}

type CreateNewRecordOutput struct {
	ID        string
	Timestamp time.Time `db:"timestamp"`
}
