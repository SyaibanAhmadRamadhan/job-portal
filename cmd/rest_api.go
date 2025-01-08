package main

import (
	"context"
	"fmt"
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/conf"
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
		fmt.Println(c)
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		go func() {
			//TODO: running server
		}()
		<-ctx.Done()
		log.Info().Msg("Received shutdown signal, shutting down server gracefully...")

		log.Info().Msg("Shutdown complete. Exiting.")
		return
	},
}
