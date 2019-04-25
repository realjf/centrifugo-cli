package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.PersistentFlags().StringVarP(&Channel, "channel", "c", "", "channel key")
}

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "allows to get channel history information",
	Long:  `allows to get channel history information (list of last messages published into channel)`,
	Run: func(cmd *cobra.Command, args []string) {
		data := params{
			ID:     UserID,
			Method: "history",
			Params: map[string]interface{}{
				"channel": Channel,
			},
		}
		WebSocket([]params{data}, nil)
		//Request("POST", "/api", []params{data}, nil)
	},
}
