package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(publishCmd)
	publishCmd.PersistentFlags().StringArrayVarP(&Channels, "channels", "cs", []string{}, "channels key")
	publishCmd.PersistentFlags().StringArrayVarP(&Data, "data", "d", []string{}, "messages to be sent")
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "allows to publish data into channel",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("publish")
	},
}
