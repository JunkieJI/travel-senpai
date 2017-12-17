package cmd

import (
	"github.com/JunkieJI/travel-senpai/api"
	"github.com/spf13/cobra"
)

// serveCmd : represents the 'start' command.
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the server.",
	Run: func(cmd *cobra.Command, args []string) {
		api.Serve()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
