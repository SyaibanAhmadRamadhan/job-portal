package handler_test

import (
	"encoding/json"
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/api"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/presentations/restfull_api/handler"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/company"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"testing"
)

func Test_handler_V1GetListCompany(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockCompanyService := company.NewMockCompanyService(mock)
	route := "/api/v1/company"
	app := fiber.New()
	handler.New(app, handler.Options{
		CompanyService: mockCompanyService,
	})

	t.Run("should be return correct", func(t *testing.T) {
		expectedPaginationInput := primitive.PaginationInput{
			Page:     1,
			PageSize: 10,
		}
		expectedPaginationOutput := primitive.CreatePaginationOutput(expectedPaginationInput, 2)
		expectedResp := api.V1GetListCompanyResponse200{
			Data: []api.V1GetListCompanyItemResponse200{
				{
					Id:   uuid.NewString(),
					Name: "KREDITKRU",
				},
				{
					Id:   uuid.NewString(),
					Name: "Goto Group",
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

		req, err := http.NewRequest("GET", route+"?page=1&page_size=10", nil)
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		mockCompanyService.EXPECT().
			GetListCompany(req.Context(), company.GetListCompanyInput{
				Pagination: expectedPaginationInput,
			}).
			Return(company.GetListCompanyOutput{
				Pagination: expectedPaginationOutput,
				Items: []company.GetListCompanyItemOutput{
					{
						ID:   expectedResp.Data[0].Id,
						Name: expectedResp.Data[0].Name,
					},
					{
						ID:   expectedResp.Data[1].Id,
						Name: expectedResp.Data[1].Name,
					},
				},
			}, nil)

		resp, err := app.Test(req, -1)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.NotNil(t, body)
		require.Equal(t, expectedRespByte, body)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
