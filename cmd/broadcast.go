package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(broadcastCmd)
	broadcastCmd.PersistentFlags().StringArrayVarP(&Channels, "channels", "cs", []string{}, "channels key")
	broadcastCmd.PersistentFlags().StringArrayVarP(&Data, "data", "d", []string{}, "messages to be sent")
}

var broadcastCmd = &cobra.Command{
	Use:   "broadcast",
	Short: "similar to publish but allows to send the same data into many channels",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
