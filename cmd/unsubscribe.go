package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(unsubscribeCmd)
	unsubscribeCmd.PersistentFlags().StringVarP(&Channel, "channel", "c", "", "channel key")
	unsubscribeCmd.PersistentFlags().StringVarP(&UserID, "user", "u", "", "user id")
}

var unsubscribeCmd = &cobra.Command{
	Use:   "unsubscribe",
	Short: "allows to unsubscribe user from channel",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
