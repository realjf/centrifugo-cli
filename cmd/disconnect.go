package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(disconnectCmd)
	disconnectCmd.PersistentFlags().StringVarP(&UserID, "user", "u", "", "user id key")
}

var disconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "allows to disconnect user by ID",
	Run: func(cmd *cobra.Command, args []string) {
		data := params{
			Method: "disconnect",
			Params: map[string]interface{}{
				"user": UserID,
			},
		}
		Request("POST", "/api", []params{data}, nil)
	},
}
