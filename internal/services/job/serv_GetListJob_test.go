package job_test

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/index_jobs"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/entity"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/job"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
	"time"
)

func Test_service_GetListJob(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockIndexJobRepository := index_jobs.NewMockJobRepository(mock)
	ctx := context.TODO()
	s := job.New(job.Options{
		IndexJobRepository: mockIndexJobRepository,
	})

	t.Run("should be return correct", func(t *testing.T) {
		expectedInput := job.GetListJobInput{
			Pagination: primitive.PaginationInput{
				Page:     1,
				PageSize: 10,
			},
			SearchKeyword: null.StringFrom("redikru"),
			CompanyID:     null.StringFrom(uuid.NewString()),
		}
		expectedOutput := job.GetListJobOutput{
			Items: []job.GetListJobItemOutput{
				{
					Company: job.GetListJobItemCompanyOutput{
						ID:   uuid.NewString(),
						Name: "KREDITKRU",
					},
					Description: "Kami mencari kandidat untuk posisi Manajer Keuangan yang memiliki pengalaman dalam industri fintech.",
					ID:          uuid.NewString(),
					Timestamp:   time.Now().UTC(),
					Title:       "Manajer Keuangan di KREDITKRU",
				},
				{
					Company: job.GetListJobItemCompanyOutput{
						ID:   uuid.NewString(),
						Name: "Goto Group",
					},
					Description: "Posisi Developer Frontend untuk membangun aplikasi inovatif di perusahaan teknologi terkemuka.",
					ID:          uuid.NewString(),
					Timestamp:   time.Now().UTC(),
					Title:       "Frontend Developer di Goto Group",
				},
			},
			Pagination: primitive.CreatePaginationOutput(expectedInput.Pagination, 2),
		}

		mockIndexJobRepository.EXPECT().
			GetAll(ctx, index_jobs.GetAllInput{
				Pagination:    expectedInput.Pagination,
				SearchKeyword: expectedInput.SearchKeyword,
				CompanyID:     expectedInput.CompanyID,
			}).
			Return(index_jobs.GetAllOutput{
				Pagination: expectedOutput.Pagination,
				Items: []entity.IndexJobEntity{
					{
						Company: entity.IndexJobCompany{
							ID:   expectedOutput.Items[0].Company.ID,
							Name: expectedOutput.Items[0].Company.Name,
						},
						Description: expectedOutput.Items[0].Description,
						ID:          expectedOutput.Items[0].ID,
						Timestamp:   expectedOutput.Items[0].Timestamp,
						Title:       expectedOutput.Items[0].Title,
					},
					{
						Company: entity.IndexJobCompany{
							ID:   expectedOutput.Items[1].Company.ID,
							Name: expectedOutput.Items[1].Company.Name,
						},
						Description: expectedOutput.Items[1].Description,
						ID:          expectedOutput.Items[1].ID,
						Timestamp:   expectedOutput.Items[1].Timestamp,
						Title:       expectedOutput.Items[1].Title,
					},
				},
			}, nil)

		output, err := s.GetListJob(ctx, expectedInput)
		require.NoError(t, err)
		require.Equal(t, expectedOutput, output)
	})
}
