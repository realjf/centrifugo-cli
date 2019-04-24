package cmd

import (
	"encoding/json"
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
		data := map[string]interface{}{
			"method": "info",
			"params": struct {
			}{},
		}
		dataJsonStr, err := json.Marshal(data)
		if err != nil {
			logrus.Error(err)
		}
		Request("POST", "/api", string(dataJsonStr), nil)
	},
}
