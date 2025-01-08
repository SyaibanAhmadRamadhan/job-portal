package index_jobs

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
)

func (r *repository) GetAll(ctx context.Context, input GetAllInput) (output GetAllOutput, err error) {
	return
}

type GetAllInput struct {
	Pagination primitive.PaginationInput
}

type GetAllOutput struct {
	Pagination primitive.PaginationOutput
	Items      []entity.IndexJobEntity
}
