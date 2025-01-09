package handler

import (
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/api"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/job"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util/primitive"
	"github.com/gofiber/fiber/v2"
	"github.com/guregu/null/v5"
	"net/http"
)

func (h *handler) V1GetListJob(c *fiber.Ctx) (err error) {
	params := api.V1GetListJobParams{}
	if err = h.httpHelper.BindQuery(c, &params); err != nil {
		return h.httpHelper.ErrResp(c, err)
	}

	jobOutput, err := h.jobService.GetListJob(c.UserContext(), job.GetListJobInput{
		Pagination: primitive.PaginationInput{
			Page:     int(params.Page),
			PageSize: int(params.PageSize),
		},
		SearchKeyword: null.StringFromPtr(params.SearchKeyword),
		CompanyID:     null.StringFromPtr(params.CompanyId),
	})
	if err != nil {
		return h.httpHelper.ErrResp(c, err)
	}

	resp := api.V1GetListJobResponse200{
		Data:       make([]api.V1GetListJobItemResponse200, 0),
		Pagination: h.httpHelper.BindToPaginationResponse(jobOutput.Pagination),
	}

	for _, item := range jobOutput.Items {
		resp.Data = append(resp.Data, api.V1GetListJobItemResponse200{
			Company: api.V1GetListJobItemCompanyResponse200{
				Id:   item.Company.ID,
				Name: item.Company.Name,
			},
			Description: item.Description,
			Id:          item.ID,
			Timestamp:   item.Timestamp,
			Title:       item.Title,
		})
	}

	return c.Status(http.StatusOK).JSON(resp)
}
