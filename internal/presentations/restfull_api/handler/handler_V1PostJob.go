package handler

import (
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/api"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/job"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *handler) V1PostJob(c *fiber.Ctx) (err error) {
	req := api.V1PostJobRequestBody{}
	if err = h.httpHelper.Bind(c, &req); err != nil {
		return h.httpHelper.ErrResp(c, err)
	}

	createJobOutput, err := h.jobService.CreateJob(c.UserContext(), job.CreateJobInput{
		CompanyName: req.Company.Name,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return h.httpHelper.ErrResp(c, err)
	}

	resp := api.V1PostJobResponse201{
		CompanyId: createJobOutput.CompanyID,
		JobId:     createJobOutput.JobID,
	}

	return c.Status(http.StatusOK).JSON(resp)
}
