package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)

var consumerEtlJobPostCmd = &cobra.Command{
	Use:   "consumer-etl-job-post",
	Short: "run consumer job post etl",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		go func() {
			//TODO: running consumer
		}()
		<-ctx.Done()
		log.Info().Msg("Received shutdown signal rest api, shutting down server gracefully...")

		log.Info().Msg("Shutdown complete. Exiting.")
		return
	},
}
