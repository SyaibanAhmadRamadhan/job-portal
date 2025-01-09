package handler

import (
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/presentations"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/company"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/job"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	validate   *presentations.Validator
	app        *fiber.App
	httpHelper *presentations.HTTPHelper

	jobService     job.JobService
	companyService company.CompanyService
}

type Options struct {
	JobService     job.JobService
	CompanyService company.CompanyService
}

func New(app *fiber.App, opts Options) {
	validate := presentations.NewValidate()
	httpHelper := presentations.NewHTTPHelper(validate.Translator, validate.Validator)
	h := handler{
		jobService:     opts.JobService,
		companyService: opts.CompanyService,
		validate:       validate,
		app:            app,
		httpHelper:     httpHelper,
	}

	h.initHandler()
}

func (h *handler) initHandler() {
	h.app.Post("/api/v1/job", h.V1PostJob)
	h.app.Get("/api/v1/job", h.V1GetListJob)

	h.app.Get("/api/v1/company", h.V1GetListCompany)
}
