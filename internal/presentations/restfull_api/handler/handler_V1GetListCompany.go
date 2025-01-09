package handler

import (
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/api"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/company"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *handler) V1GetListCompany(c *fiber.Ctx) (err error) {
	paginationInput := h.httpHelper.BindToPaginationInput(c)

	companyOutput, err := h.companyService.GetListCompany(c.UserContext(), company.GetListCompanyInput{
		Pagination: paginationInput,
	})
	if err != nil {
		return h.httpHelper.ErrResp(c, err)
	}

	resp := api.V1GetListCompanyResponse200{
		Data:       make([]api.V1GetListCompanyItemResponse200, 0),
		Pagination: h.httpHelper.BindToPaginationResponse(companyOutput.Pagination),
	}
	for _, item := range companyOutput.Items {
		resp.Data = append(resp.Data, api.V1GetListCompanyItemResponse200{
			Id:   item.ID,
			Name: item.Name,
		})
	}

	return c.Status(http.StatusOK).JSON(resp)
}
