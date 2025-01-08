package main

import (
	"github.com/SyaibanAhmadRamadhan/job-portal/internal/util"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{}

	rootCmd.AddCommand(restApiCmd)
	rootCmd.AddCommand(consumerEtlJobPostCmd)

	err := rootCmd.Execute()
	util.Panic(err)
}
