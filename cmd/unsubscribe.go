package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(unsubscribeCmd)
	unsubscribeCmd.PersistentFlags().StringVarP(&Channel, "channel", "c", "", "channel key")
	unsubscribeCmd.PersistentFlags().Uint32VarP(&UserID, "user", "u", 0, "user id")
}

var unsubscribeCmd = &cobra.Command{
	Use:   "unsubscribe",
	Short: "allows to unsubscribe user from channel",
	Run: func(cmd *cobra.Command, args []string) {
		data := params{
			ID:     UserID,
			Method: "unsubscribe",
			Params: map[string]interface{}{
				"channel": Channel,
			},
		}
		Request("POST", "/api", []params{data}, nil)
	},
}
