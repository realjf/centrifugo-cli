package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(presenceCmd)
}

// 在线情况
var presenceCmd = &cobra.Command{
	Use:   "presence",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("presence")
	},
}
