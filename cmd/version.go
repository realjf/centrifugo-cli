package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of centrifugo-cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("centrifugo-cli version 1.0.0")
	},
}
