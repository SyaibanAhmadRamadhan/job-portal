package main

import (
	"context"
	"fmt"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/infra"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/presentations/restfull_api"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/companies"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/jobs"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/eventbus"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/company"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/job"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)

var restApiCmd = &cobra.Command{
	Use:   "rest-api",
	Short: "run rest api",
	Run: func(cmd *cobra.Command, args []string) {
		c := conf.LoadConfig()
		closeFnOtel := infra.NewOtel(&c.Otel, c.AppName)
		db, closeFnDB := infra.NewPostgreCommand(&c.Database)
		//dbQuery, closeFnDBQuery := infra.NewPostgreQuery(&c.Database)
		wsqlxDb := wsqlx.NewRdbms(db)
		//wsqlxQuery := wsqlx.NewRdbms(dbQuery)
		kafkaBroker, closeFnKafkaBroker := infra.NewKafkaWriter(c.Kafka)
		redisClient, closeFnRedis := infra.NewRedisWithOtel(&c.Redis, c.AppName)

		// REPO LAYER
		jobsRepository := jobs.New(wsqlxDb)
		//indexJobsRepository := index_jobs.New(wsqlxQuery)
		companiesRepository := companies.New(wsqlxDb, redisClient)
		eventbusRepository := eventbus.New(kafkaBroker)

		// SERVICE LAYER
		jobService := job.New(job.Options{
			IndexJobRepository:          nil,
			JobRepository:               jobsRepository,
			CompanyRepository:           companiesRepository,
			EventBusPublisherRepository: eventbusRepository,
			DBTx:                        wsqlxDb,
		})
		companyService := company.New(company.Options{
			CompanyRepository: companiesRepository,
		})

		server := restfull_api.New(restfull_api.Presenter{
			AppPort: c.Port,
			AppName: c.AppName,
			Dependency: restfull_api.Dependency{
				JobService:     jobService,
				CompanyService: companyService,
			},
		})

		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		go func() {
			if err := server.Listen(fmt.Sprintf(":%d", c.Port)); err != nil {
				log.Error().Err(err).Msg("failed listen server")
				stop()
			}
		}()

		<-ctx.Done()
		log.Info().Msg("Received shutdown signal, shutting down server gracefully...")

		if err := closeFnDB(context.TODO()); err != nil {
			log.Error().Err(err).Msgf("failed closed db command: %v", err)
		}

		//if err := closeFnDBQuery(context.TODO()); err != nil {
		//	log.Error().Err(err).Msgf("failed closed db query: %v", err)
		//}

		if err := closeFnOtel(context.TODO()); err != nil {
			log.Error().Err(err).Msgf("failed closed otel: %v", err)
		}

		if err := closeFnKafkaBroker(context.TODO()); err != nil {
			log.Error().Err(err).Msgf("failed closed kafka broker: %v", err)
		}

		if err := closeFnRedis(context.TODO()); err != nil {
			log.Error().Err(err).Msgf("failed closed redis client: %v", err)
		}

		log.Info().Msg("Shutdown complete. Exiting.")
		return
	},
}
