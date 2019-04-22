package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"net/url"
)

func init() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.PersistentFlags().StringVarP(&Address, "address", "a", "", "bind your centrifugo to specific interface address")
	connectCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8000, "port to bind centrifugo to")
	connectCmd.PersistentFlags().StringVarP(&Engine, "engine", "e", "memory", "engine to use - memory or redis")
}

func GetHost() string {
	return fmt.Sprintf("%s", Address)
}

func GetPath() string {
	return fmt.Sprintf("/connection")
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "connect to centrifugo server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Infof("address：%s", Address)
		logrus.Infof("port：%d", Port)
		logrus.Infof("engine：%s", Engine)
		logrus.Infoln("connecting...")
		logrus.Infof("token：%s", _token())
		client := http.Client{}
		var req = new(http.Request)
		req.Header = http.Header{}
		req.Header.Set("Authentication", "token "+_token())
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Connection", "keep-alive")
		req.URL = new(url.URL)
		req.URL.Host = Address
		req.URL.Path = GetPath()
		resp, err := client.Do(req)
		if err != nil {
			logrus.Error(err)
		}
		response, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			logrus.Error(err)
		}
		logrus.Infof("%s", string(response))
	},
}
