package restfull_api

import (
	"github.com/SyaibanAhmadRamadhan/job-portal/generated/api"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/presentations/restfull_api/handler"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/company"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/job"
	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"net/http"
	"time"
)

type Presenter struct {
	AppPort    int
	AppName    string
	Dependency Dependency
}

type Dependency struct {
	JobService     job.JobService
	CompanyService company.CompanyService
}

func New(presenter Presenter) *fiber.App {
	app := fiber.New()
	initMiddleware(app, presenter.AppName, presenter.AppPort)

	handler.New(app, handler.Options{
		JobService:     presenter.Dependency.JobService,
		CompanyService: presenter.Dependency.CompanyService,
	})
	return app
}

func initMiddleware(app *fiber.App, appName string, appPort int) {
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(limiter.New(limiter.Config{
		Max:        10,
		Expiration: 1 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(http.StatusTooManyRequests).JSON(api.Error{
				Message: "too many requests",
			})
		},
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3002",
		AllowHeaders:     "Origin, Content-Type, Accept, X-User-Id, X-Request-Id, X-Correlation-Id, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Use(otelfiber.Middleware(
		otelfiber.WithServerName(appName),
		otelfiber.WithPort(appPort),
		otelfiber.WithCollectClientIP(true),
	))
}
