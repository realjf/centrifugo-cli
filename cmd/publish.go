package cmd

import (
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
		Request("GET", "/api", "", nil)
	},
}
