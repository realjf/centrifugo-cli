package cmd

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.PersistentFlags().StringVarP(&Address, "address", "a", "", "bind your centrifugo to specific interface address")
	connectCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8000, "port to bind centrifugo to")
	connectCmd.PersistentFlags().StringVarP(&Engine, "engine", "e", "memory", "engine to use - memory or redis")
	connectCmd.PersistentFlags().Uint32VarP(&UserID, "user", "u", 0, "user id")
	connectCmd.PersistentFlags().StringVarP(&Path, "path", "l", "/connection/websocket", "websocket path to bind")
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to centrifugo server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("address：%s", Address)
		logrus.Infof("port：%d", Port)
		logrus.Infof("engine：%s", Engine)
		logrus.Infof("user id: %d", UserID)
		if UserID <= 0 {
			logrus.Error("need user id to continue")
		}
		logrus.Infoln("connecting...")
		data := params{
			ID:     UserID,
			Method: "connect",
			Params: map[string]interface{}{
				"token": _token(UserID),
				"data": struct {
				}{},
			},
		}
		message := WebSocket([]params{data}, nil)
		var result struct {
			ID     int `json:"id"`
			Result struct {
				Client  string `json:"client"`
				Version string `json:"version"`
				Expires bool   `json:"expires"`
				Ttl     int    `json:"ttl"`
			}
		}
		err := json.Unmarshal(message.([]byte), &result)
		if err != nil {
			logrus.Error(err)
		}
		ClientConnectionID = result.Result.Client

		//Request("POST", "/api", []params{data}, nil)
	},
}
