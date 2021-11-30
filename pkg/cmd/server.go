package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ryota-sakamoto/topsy-turvy/pkg/server"
)

const (
	ServerIDFlag = "server-id"
)

func init() {
	serverCmd.Flags().String(ServerIDFlag, "", "Server ID")
	serverCmd.MarkFlagRequired(ServerIDFlag)
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	RunE: func(cmd *cobra.Command, args []string) error {
		serverID, err := cmd.Flags().GetString(ServerIDFlag)
		if err != nil {
			return err
		}

		return server.Start(serverID)
	},
}
