package cmd

import (
	"log"

	"github.com/JunkieJI/travel-senpai/config"
	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd : the base command when called without any subcommands.
var RootCmd = &cobra.Command{
	Use:   "travel-senpai",
	Short: "Travel Senpai",
}

func init() {
	// Read in a config file using viper.
	cobra.OnInitialize(config.Init)

	// Initialize persistent and local flags if required.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config.yml", "Path to the configuration file.")

}

// Execute : called by main.main().
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
