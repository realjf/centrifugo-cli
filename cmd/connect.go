package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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

func Request(method string, path string, data string) {
	client := http.Client{}
	uri := url.URL{}
	uri.Host = GetHost()
	uri.Scheme = "http"
	uri.Path = path
	req, err := http.NewRequest(method, uri.String(), strings.NewReader(data))
	if err != nil {
		logrus.Error(err)
	}
	req.Header.Set("Authentication", "token "+_token())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")

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
		Request("GET", GetPath(), "")
	},
}
