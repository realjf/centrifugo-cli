package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(presenceStatsCmd)
	presenceStatsCmd.PersistentFlags().StringVarP(&Channel, "channel", "c", "", "channel key")
}

var presenceStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "allows to get short channel presence information",
	Run: func(cmd *cobra.Command, args []string) {
		data := params{
			ID:     UserID,
			Method: "presence_stats",
			Params: map[string]interface{}{
				"channel": Channel,
			},
		}

		WebSocket([]params{data}, nil)
		//Request("GET", "/api", []params{data}, nil)
	},
}
