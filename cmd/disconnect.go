package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(disconnectCmd)
	disconnectCmd.PersistentFlags().Uint32VarP(&UserID, "user", "u", 0, "user id")
}

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "allows to disconnect user by ID",
	Run: func(cmd *cobra.Command, args []string) {
		data := params{
			ID:     UserID,
			Method: "disconnect",
			Params: map[string]interface{}{
				"user": UserID,
			},
		}
		WebSocket([]params{data}, nil)
		//Request("POST", "/api", []params{data}, nil)
	},
}
