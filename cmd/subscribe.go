package cmd

import (
	"github.com/spf13/cobra"
	"net/http"
)

func init() {
	rootCmd.AddCommand(subscribeCmd)
	subscribeCmd.PersistentFlags().StringVarP(&Channel, "channel", "c", "", "channel key")
	subscribeCmd.PersistentFlags().Uint32VarP(&UserID, "user", "u", 0, "user id")
}

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "subscribe on channels",
	Run: func(cmd *cobra.Command, args []string) {
		data := params{
			ID:     UserID,
			Method: "subscribe",
			Params: map[string]interface{}{
				"channel": Channel,
				"token":   _token(UserID),
				"client":  ClientConnectionID,
			},
		}

		header := http.Header{}
		header.Add("Authorization", "token "+_token(UserID))
		WebSocket([]params{data}, header)
		//Request("POST", "/api", []params{data}, nil)
	},
}
