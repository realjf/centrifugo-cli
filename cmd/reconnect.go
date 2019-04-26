package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(reconnect)
}

var reconnect = &cobra.Command{
	Use:   "reconnect",
	Short: "reconnect",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infoln("reconnecting...")
		WebSocketConn.WriteClose(1)
		WebSocketConn.Close()
		data := params{
			ID:     UserID,
			Method: "connect",
			Params: map[string]interface{}{
				"token": _token(UserID),
				"data": struct {
				}{},
			},
		}
		WebSocket([]params{data}, nil)
	},
}
