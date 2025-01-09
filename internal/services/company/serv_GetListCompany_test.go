package company_test

import (
	"context"
	"errors"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/company"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/tracer"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func Test_service_GetListCompany(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()
	ctx := context.TODO()
	mockCompanyRepository := companies.NewMockCompanyRepository(mock)

	s := company.New(company.Options{
		CompanyRepository: mockCompanyRepository,
	})

	t.Run("should be return correct is cache", func(t *testing.T) {
		expectedInput := company.GetListCompanyInput{
			Pagination: primitive.PaginationInput{
				Page:     1,
				PageSize: 10,
			},
		}
		expectedOutput := company.GetListCompanyOutput{
			Pagination: primitive.CreatePaginationOutput(expectedInput.Pagination, 2),
			Items: []company.GetListCompanyItemOutput{
				{
					ID:   uuid.NewString(),
					Name: "KREDITKRU",
				},
				{
					ID:   uuid.NewString(),
					Name: "Goto Group",
				},
			},
		}

		mockCompanyRepository.EXPECT().
			GetAllCache(ctx, companies.GetAllCacheInput{
				Pagination: expectedInput.Pagination,
			}).
			Return(companies.GetAllCacheOutput{
				Data: entity.CompanyCacheEntity{
					Pagination: expectedOutput.Pagination,
					Items: []entity.CompanyEntity{
						{
							ID:   expectedOutput.Items[0].ID,
							Name: expectedOutput.Items[0].Name,
						},
						{
							ID:   expectedOutput.Items[1].ID,
							Name: expectedOutput.Items[1].Name,
						},
					},
				},
			}, nil)

		output, err := s.GetListCompany(ctx, expectedInput)
		require.NoError(t, err)
		require.Equal(t, expectedOutput, output)
	})

	t.Run("should be return correct is no cache", func(t *testing.T) {
		expectedInput := company.GetListCompanyInput{
			Pagination: primitive.PaginationInput{
				Page:     1,
				PageSize: 10,
			},
		}
		expectedOutput := company.GetListCompanyOutput{
			Pagination: primitive.CreatePaginationOutput(expectedInput.Pagination, 2),
			Items: []company.GetListCompanyItemOutput{
				{
					ID:   uuid.NewString(),
					Name: "KREDITKRU",
				},
				{
					ID:   uuid.NewString(),
					Name: "Goto Group",
				},
			},
		}

		mockCompanyRepository.EXPECT().
			GetAllCache(ctx, companies.GetAllCacheInput{
				Pagination: expectedInput.Pagination,
			}).
			Return(companies.GetAllCacheOutput{}, tracer.ErrNotFound(errors.New("data is nil")))

		mockCompanyRepository.EXPECT().
			GetAll(ctx, companies.GetAllInput{
				Pagination: expectedInput.Pagination,
			}).
			Return(companies.GetAllOutput{
				Pagination: expectedOutput.Pagination,
				Items: []entity.CompanyEntity{
					{
						ID:   expectedOutput.Items[0].ID,
						Name: expectedOutput.Items[0].Name,
					},
					{
						ID:   expectedOutput.Items[1].ID,
						Name: expectedOutput.Items[1].Name,
					},
				},
			}, nil)

		mockCompanyRepository.EXPECT().
			SetCache(ctx, gomock.Any()).Return(nil)

		output, err := s.GetListCompany(ctx, expectedInput)
		require.NoError(t, err)
		require.Equal(t, expectedOutput, output)
	})
}
