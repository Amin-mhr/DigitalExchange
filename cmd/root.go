/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"DigitalExchange/cmd/app"
	"DigitalExchange/cmd/database"
	"DigitalExchange/src/config"
	"github.com/spf13/cobra"
	"log"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "DigitalExchange",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(
		app.AppCmd,
		database.DatabaseCmd,
		//cron.CronCmd,
	)
}

func initConfig() {
	// Initialize configuration
	err := config.Init()
	if err != nil {
		log.Fatalf("config Service: Failed to Initialize. %v", err)
	}
	log.Println("config Service: Initialized Successfully.")
}
