package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(publishCmd)
	publishCmd.PersistentFlags().StringVarP(&Channel, "channel", "c", "", "channel key")
	publishCmd.PersistentFlags().StringVarP(&Data, "data", "d", "", "messages to be sent")
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "allows to publish data into channel",
	Long:  `publish -c=channel#userId -d=data`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infoln("publishing...")
		data := params{
			Method: "publish",
			Params: map[string]interface{}{
				"channel": Channel,
				"data": map[string]string{
					"text": Data,
				},
			},
		}
		logrus.Infof("channel：%v", Channel)
		Request("POST", "/api", []params{data}, nil)
	},
}
