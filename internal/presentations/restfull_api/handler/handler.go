package handler

import (
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/presentations"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/job"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	validate   *presentations.Validator
	app        *fiber.App
	httpHelper *presentations.HTTPHelper

	jobService job.JobService
}

type Options struct {
	JobService job.JobService
}

func New(app *fiber.App, opts Options) {
	validate := presentations.NewValidate()
	httpHelper := presentations.NewHTTPHelper(validate.Translator, validate.Validator)
	h := handler{
		jobService: opts.JobService,
		validate:   validate,
		app:        app,
		httpHelper: httpHelper,
	}

	app.Post("/api/v1/job", h.V1PostJob)

}
