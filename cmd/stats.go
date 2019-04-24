package cmd

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
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
		data := map[string]interface{}{
			"method": "presence_stats",
			"params": map[string]interface{}{
				"channel": Channel,
			},
		}
		dataJsonStr, err := json.Marshal(data)
		if err != nil {
			logrus.Error(err)
		}
		Request("GET", "/api", string(dataJsonStr), nil)
	},
}
