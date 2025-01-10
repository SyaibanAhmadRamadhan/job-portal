package main

import (
	"context"
	ekafka "github.com/SyaibanAhmadRamadhan/event-bus/kafka"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/infra"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/datastore/index_jobs"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/repositories/eventbus"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/services/job"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)

var consumerEtlJobPostCmd = &cobra.Command{
	Use:   "consumer-etl-job-post",
	Short: "run consumer job post etl",
	Run: func(cmd *cobra.Command, args []string) {
		c := conf.LoadConfig()
		closeFnOtel := infra.NewOtel(&c.Otel, c.AppName)
		kafkaReader := ekafka.New(ekafka.WithOtel())
		esClient, closeFnEs := infra.NewES(&c.Database.ElasticsearchConfig)

		// REPO LAYER
		indexJobsRepository := index_jobs.New(esClient)
		eventbusRepository := eventbus.New(kafkaReader, &c.Kafka)

		// SERVICE LAYER
		jobService := job.New(job.Options{
			IndexJobRepository:         indexJobsRepository,
			EventBusConsumerRepository: eventbusRepository,
		})

		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		go func() {
			if err := jobService.ConsumerPostJobETL(ctx); err != nil {
				log.Error().Err(err).Msg("failed consumer post job etl")
				stop()
			}
		}()
		<-ctx.Done()
		log.Info().Msg("Received shutdown signal rest api, shutting down server gracefully...")

		if err := closeFnEs(context.TODO()); err != nil {
			log.Error().Err(err).Msgf("failed closed redis client: %v", err)
		}

		if err := closeFnOtel(context.TODO()); err != nil {
			log.Error().Err(err).Msgf("failed closed otel: %v", err)
		}

		log.Info().Msg("Shutdown complete. Exiting.")
		return
	},
}
