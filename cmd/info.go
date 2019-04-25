package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "allows to get information about running Centrifugo nodes",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infoln("requesting...")
		data := params{
			ID:     UserID,
			Method: "info",
			Params: map[string]interface{}{},
		}

		WebSocket([]params{data}, nil)
		//Request("POST", "/api", []params{data}, nil)
	},
}
