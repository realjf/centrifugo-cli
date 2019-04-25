package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(broadcastCmd)
	broadcastCmd.PersistentFlags().StringArrayVarP(&Channels, "channels", "c", []string{}, "channels key")
	broadcastCmd.PersistentFlags().StringVarP(&Data, "data", "d", "", "messages to be sent")
}

var broadcastCmd = &cobra.Command{
	Use:   "broadcast",
	Short: "similar to publish but allows to send the same data into many channels",
	Run: func(cmd *cobra.Command, args []string) {
		data := params{
			ID:     UserID,
			Method: "broadcast",
			Params: map[string]interface{}{
				"channels": Channels,
				"data": map[string]string{
					"text": Data,
				},
			},
		}
		WebSocket([]params{data}, nil)
		//Request("POST", "/api", []params{data}, nil)
	},
}
