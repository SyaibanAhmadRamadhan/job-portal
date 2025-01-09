package companies

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/jmoiron/sqlx"
)

func (r *repository) GetAll(ctx context.Context, input GetAllInput) (output GetAllOutput, err error) {
	query := r.sq.Select("id", "name").From("companies")
	queryCount := r.sq.Select("count(*)").From("companies")

	output = GetAllOutput{
		Items: make([]entity.CompanyEntity, 0),
	}

	paginationOutput, err := r.rdbms.QuerySqPagination(ctx, queryCount, query, wsqlx.PaginationInput{
		Page:     int64(input.Pagination.Page),
		PageSize: int64(input.Pagination.PageSize),
	}, func(rows *sqlx.Rows) (err error) {
		item := entity.CompanyEntity{}

		for rows.Next() {
			if err = rows.StructScan(&item); err != nil {
				return tracer.ErrInternalServer(err)
			}
			output.Items = append(output.Items, item)
		}

		return
	})
	if err != nil {
		return output, tracer.ErrInternalServer(err)
	}

	output.Pagination = primitive.PaginationOutput{
		Page:      int(paginationOutput.Page),
		PageSize:  int(paginationOutput.PageSize),
		PageCount: int(paginationOutput.PageCount),
		TotalData: int(paginationOutput.TotalData),
	}
	return
}

type GetAllInput struct {
	Pagination primitive.PaginationInput
}

type GetAllOutput struct {
	Pagination primitive.PaginationOutput
	Items      []entity.CompanyEntity
}
