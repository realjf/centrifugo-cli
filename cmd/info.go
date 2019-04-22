package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show centrifugo server information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info")
	},
}
