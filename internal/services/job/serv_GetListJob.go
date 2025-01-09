package job

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/index_jobs"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	"github.com/guregu/null/v5"
	"time"
)

func (s *service) GetListJob(ctx context.Context, input GetListJobInput) (output GetListJobOutput, err error) {
	getAllJobOutput, err := s.indexJobRepository.GetAll(ctx, index_jobs.GetAllInput{
		Pagination:    input.Pagination,
		SearchKeyword: input.SearchKeyword,
		CompanyID:     input.CompanyID,
	})
	if err != nil {
		return output, tracer.ErrInternalServer(err)
	}

	output = GetListJobOutput{
		Pagination: getAllJobOutput.Pagination,
		Items:      make([]GetListJobItemOutput, 0),
	}
	for _, item := range getAllJobOutput.Items {
		output.Items = append(output.Items, GetListJobItemOutput{
			Company: GetListJobItemCompanyOutput{
				ID:   item.Company.ID,
				Name: item.Company.Name,
			},
			ID:          item.ID,
			Title:       item.Title,
			Description: item.Description,
			Timestamp:   item.Timestamp,
		})
	}
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
	Timestamp   time.Time
}

type GetListJobItemCompanyOutput struct {
	ID   string
	Name string
}
