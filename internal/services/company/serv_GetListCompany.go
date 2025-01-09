package company

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	"time"
)

func (s *service) GetListCompany(ctx context.Context, input GetListCompanyInput) (output GetListCompanyOutput, err error) {
	isNoCache := false

	companiesCacheOutput, err := s.companyRepository.GetAllCache(ctx, companies.GetAllCacheInput{
		Pagination: input.Pagination,
	})
	if err != nil {
		if !tracer.ErrorAsNotFound(err) {
			return output, tracer.StackTrace(err)
		}
		err = nil
		isNoCache = true
	}

	if isNoCache {
		companiesOutput, err := s.companyRepository.GetAll(ctx, companies.GetAllInput{
			Pagination: input.Pagination,
		})
		if err != nil {
			return output, tracer.StackTrace(err)
		}

		if len(companiesOutput.Items) > 0 {
			err = s.companyRepository.SetCache(ctx, companies.SetCacheInput{
				ExpiredAt: time.Now().UTC().Add(24 * time.Hour),
				Payload: entity.CompanyCacheEntity{
					Pagination: companiesOutput.Pagination,
					Items:      companiesOutput.Items,
				},
			})
		}

		companiesCacheOutput.Data.Items = companiesOutput.Items
		companiesCacheOutput.Data.Pagination = companiesOutput.Pagination
	}

	output = GetListCompanyOutput{
		Pagination: companiesCacheOutput.Data.Pagination,
		Items:      make([]GetListCompanyItemOutput, 0),
	}

	for _, item := range companiesCacheOutput.Data.Items {
		output.Items = append(output.Items, GetListCompanyItemOutput{
			ID:   item.ID,
			Name: item.Name,
		})
	}
	return
}

type GetListCompanyInput struct {
	Pagination primitive.PaginationInput
}

type GetListCompanyOutput struct {
	Pagination primitive.PaginationOutput
	Items      []GetListCompanyItemOutput
}

type GetListCompanyItemOutput struct {
	ID   string
	Name string
}
