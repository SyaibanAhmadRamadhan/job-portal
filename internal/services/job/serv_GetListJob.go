package job

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/guregu/null/v5"
)

func (s *service) GetListJob(ctx context.Context, input GetListJobInput) (output GetListJobOutput, err error) {
	return
}

type GetListJobInput struct {
	Pagination    primitive.PaginationInput
	SearchKeyword null.String
	CompanyID     null.String
}

type GetListJobOutput struct {
	Pagination primitive.PaginationOutput
	Items      []GetListJobItemOutput
}

type GetListJobItemOutput struct {
	Company     GetListJobItemCompanyOutput
	ID          string
	Title       string
	Description string
	Timestamp   string
}

type GetListJobItemCompanyOutput struct {
	ID   string
	Name string
}
