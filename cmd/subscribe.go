package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(subscribeCmd)
	subscribeCmd.PersistentFlags().StringVarP(&Channel, "channel", "c", "", "channel key")
}

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "subscribe on channels",
	Run: func(cmd *cobra.Command, args []string) {
		data := params{
			Method: "subscribe",
			Params: map[string]interface{}{
				"channel": Channel,
			},
		}
		Request("POST", "/api", []params{data}, nil)
	},
}
