package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(publishCmd)
}

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "publish message to channel on centrifugo server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("publish")
	},
}
