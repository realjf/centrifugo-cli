package cmd

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read message",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("reading message...")

		_, message, err := WebSocketConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Error(err)
			}
			logrus.Error(err)
		}

		logrus.Infof("received: %s\n", message)
	},
}
