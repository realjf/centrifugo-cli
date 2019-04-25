package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Address string // centrifugo server address
var Port int       // port to bind
var Engine string

var Secret string
var ApiKey string
var Token string

var UserID uint32
var Channel string // 当前channel
var Channels []string

var Path string // websocket path

var Data string

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "show currently config information",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("address：%s", Address)
		logrus.Infof("port：%d", Port)
		logrus.Infof("engine：%s", Engine)
		logrus.Infof("secret：%s", Secret)
		logrus.Infof("ApiKey：%s", ApiKey)
		logrus.Infof("Token：%s", Token)
		logrus.Infof("UserID：%s", UserID)
		logrus.Infof("Channel：%s", Channel)
		logrus.Infof("Channels：%v", Channels)
		logrus.Infof("Data：%v", Data)
	},
}
