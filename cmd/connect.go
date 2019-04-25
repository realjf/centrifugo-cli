package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.PersistentFlags().StringVarP(&Address, "address", "a", "", "bind your centrifugo to specific interface address")
	connectCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8000, "port to bind centrifugo to")
	connectCmd.PersistentFlags().StringVarP(&Engine, "engine", "e", "memory", "engine to use - memory or redis")
	connectCmd.PersistentFlags().StringVarP(&ConnectProtocol, "connect-protocol", "c", "http", "communication protocol")
	connectCmd.PersistentFlags().StringVarP(&UserID, "user", "u", "", "user id")
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to centrifugo server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("address：%s", Address)
		logrus.Infof("port：%d", Port)
		logrus.Infof("engine：%s", Engine)
		logrus.Infoln("connecting...")
		data := params{
			Method: "connect",
			Params: map[string]interface{}{
				"token": "JWT",
				"data": struct {
				}{},
			},
		}
		logrus.Infof("user id: %s", UserID)
		Request("POST", "/api", []params{data}, nil)
	},
}
