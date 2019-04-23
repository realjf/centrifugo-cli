package cmd

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(publishCmd)
	publishCmd.PersistentFlags().StringArrayVarP(&Channels, "channels", "c", []string{}, "channels key")
	publishCmd.PersistentFlags().StringArrayVarP(&Data, "data", "d", []string{}, "messages to be sent")
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "allows to publish data into channel",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infoln("publishing...")
		data := map[string]interface{}{
			"method": "publish",
			"parmas": map[string]interface{}{
				"channel": Channel,
				"data":    Data,
			},
		}
		dataJsonStr, err := json.Marshal(data)
		if err != nil {
			logrus.Error(err)
		}
		Request("POST", "/api", string(dataJsonStr), nil)
	},
}
