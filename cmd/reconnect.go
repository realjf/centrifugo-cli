package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(reconnect)
}

var reconnect = &cobra.Command{
	Use:   "reconnect",
	Short: "reconnect",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
