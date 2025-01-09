package handler_test

import (
	"bytes"
	"encoding/json"
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/api"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/presentations/restfull_api/handler"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/job"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"testing"
)

func Test_handler_V1JobPost(t *testing.T) {
	mock := gomock.NewController(t)
	defer mock.Finish()

	mockJobService := job.NewMockJobService(mock)
	route := "/api/v1/job"
	app := fiber.New()
	handler.New(app, handler.Options{
		JobService: mockJobService,
	})

	t.Run("should be return error validation", func(t *testing.T) {
		expectedResp := api.Error400{
			Errors: map[string][]string{
				"title": {
					"title must be at least 5 characters in length",
				},
			},
			Message: "invalid your request",
		}
		expectedRespByte, err := json.Marshal(expectedResp)
		require.NoError(t, err)

		reqBody := api.V1PostJobRequestBody{
			Company: api.V1PostJobRequestBodyCompany{
				Name: "REDIKRU",
			},
			Description: "Description",
			Title:       "test",
		}
		reqBodyByte, err := json.Marshal(reqBody)
		require.NoError(t, err)
		req, err := http.NewRequest("POST", route, bytes.NewBuffer(reqBodyByte))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req, -1)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.NotNil(t, body)
		require.Equal(t, expectedRespByte, body)
		require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})

	t.Run("should be return success", func(t *testing.T) {
		expectedResp := api.V1PostJobResponse201{
			CompanyId: uuid.New().String(),
			JobId:     uuid.New().String(),
		}
		expectedRespByte, err := json.Marshal(expectedResp)
		require.NoError(t, err)

		reqBody := api.V1PostJobRequestBody{
			Company: api.V1PostJobRequestBodyCompany{
				Name: "REDIKRU",
			},
			Description: "Description",
			Title:       "backend developer",
		}
		reqBodyByte, err := json.Marshal(reqBody)
		require.NoError(t, err)

		req, err := http.NewRequest("POST", route, bytes.NewBuffer(reqBodyByte))
		require.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		mockJobService.EXPECT().
			CreateJob(req.Context(), job.CreateJobInput{
				CompanyName: reqBody.Company.Name,
				Title:       reqBody.Title,
				Description: reqBody.Description,
			}).
			Return(job.CreateJobOutput{
				CompanyID: expectedResp.CompanyId,
				JobID:     expectedResp.JobId,
			}, nil)

		resp, err := app.Test(req, -1)

		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.NotNil(t, body)
		require.Equal(t, expectedRespByte, body)
		require.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
