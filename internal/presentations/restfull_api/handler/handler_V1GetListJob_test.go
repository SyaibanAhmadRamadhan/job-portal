package handler_test

import (
	"encoding/json"
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/api"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/presentations/restfull_api/handler"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/job"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/guregu/null/v5"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"testing"
	"time"
)

func Test_handler_V1GetListJob(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockJobService := job.NewMockJobService(mock)
	route := "/api/v1/job"
	app := fiber.New()
	handler.New(app, handler.Options{
		JobService: mockJobService,
	})

	t.Run("should be return correct", func(t *testing.T) {
		expectedPaginationInput := primitive.PaginationInput{
			Page:     1,
			PageSize: 10,
		}
		expectedPaginationOutput := primitive.CreatePaginationOutput(expectedPaginationInput, 2)
		expectedResp := api.V1GetListJobResponse200{
			Data: []api.V1GetListJobItemResponse200{
				{
					Company: api.V1GetListJobItemCompanyResponse200{
						Id:   uuid.NewString(),
						Name: "KREDITKRU",
					},
					Description: "Kami mencari kandidat untuk posisi Manajer Keuangan yang memiliki pengalaman dalam industri fintech.",
					Id:          uuid.NewString(),
					Timestamp:   time.Now().UTC(),
					Title:       "Manajer Keuangan di KREDITKRU",
				},
				{
					Company: api.V1GetListJobItemCompanyResponse200{
						Id:   uuid.NewString(),
						Name: "Goto Group",
					},
					Description: "Posisi Developer Frontend untuk membangun aplikasi inovatif di perusahaan teknologi terkemuka.",
					Id:          uuid.NewString(),
					Timestamp:   time.Now().UTC(),
					Title:       "Frontend Developer di Goto Group",
				},
			},
			Pagination: api.PaginationResponse{
				Page:      1,
				PageCount: 1,
				PageSize:  10,
				TotalData: 2,
			},
		}
		expectedRespByte, err := json.Marshal(expectedResp)
		require.NoError(t, err)

		req, err := http.NewRequest("GET", route+"?page=1&page_size=10&search_keyword=redikru&company_id="+expectedResp.Data[0].Company.Id, nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		mockJobService.EXPECT().
			GetListJob(req.Context(), job.GetListJobInput{
				Pagination:    expectedPaginationInput,
				SearchKeyword: null.StringFrom("redikru"),
				CompanyID:     null.StringFrom(expectedResp.Data[0].Company.Id),
			}).
			Return(job.GetListJobOutput{
				Pagination: expectedPaginationOutput,
				Items: []job.GetListJobItemOutput{
					{
						Company: job.GetListJobItemCompanyOutput{
							ID:   expectedResp.Data[0].Company.Id,
							Name: expectedResp.Data[0].Company.Name,
						},
						Description: expectedResp.Data[0].Description,
						ID:          expectedResp.Data[0].Id,
						Timestamp:   expectedResp.Data[0].Timestamp,
						Title:       expectedResp.Data[0].Title,
					},
					{
						Company: job.GetListJobItemCompanyOutput{
							ID:   expectedResp.Data[1].Company.Id,
							Name: expectedResp.Data[1].Company.Name,
						},
						Description: expectedResp.Data[1].Description,
						ID:          expectedResp.Data[1].Id,
						Timestamp:   expectedResp.Data[1].Timestamp,
						Title:       expectedResp.Data[1].Title,
					},
				},
			}, nil)

		resp, err := app.Test(req, -1)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.NotNil(t, body)
		require.Equal(t, string(expectedRespByte), string(body))
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
