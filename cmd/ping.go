package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(pingCmd)
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "to maintain connection alive and detect broken connections",
	Run: func(cmd *cobra.Command, args []string) {
		data := params{
			Method: "ping",
			Params: map[string]interface{}{},
		}
		Request("POST", "/api", []params{data}, nil)
	},
}
