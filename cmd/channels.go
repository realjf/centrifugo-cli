package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(channelsCmd)
}

var channelsCmd = &cobra.Command{
	Use:   "channels",
	Short: "allows to get list of active (with one or more subscribers) channels",
	Run: func(cmd *cobra.Command, args []string) {
		data := params{
			Method: "channels",
			Params: map[string]interface{}{},
		}
		Request("POST", "/api", []params{data}, nil)
	},
}
