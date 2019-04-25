package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(presenceCmd)
	presenceCmd.PersistentFlags().StringVarP(&Channel, "channel", "c", "chat", "channel key")
}

// 在线情况
var presenceCmd = &cobra.Command{
	Use:   "presence",
	Short: "allows to get channel presence information",
	Long:  `allows to get channel presence information(all clients currently subscribed on this channel)`,
	Run: func(cmd *cobra.Command, args []string) {
		data := params{
			ID:     UserID,
			Method: "presence",
			Params: map[string]interface{}{
				"channel": Channel,
			},
		}
		WebSocket([]params{data}, nil)
		//Request("POST", "/api", []params{data}, nil)
	},
}
