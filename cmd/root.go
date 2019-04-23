package cmd

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&ApiKey, "api-key", "k", "", "set api key for http api")
	rootCmd.PersistentFlags().StringVarP(&Secret, "secret", "s", "178e1a29-fbc8-4422-a47e-3b66a411677b", "secret key for http api")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.SetDefault("author", "Real JF <real_jf@hotmail.com>")
	viper.SetDefault("license", "apache")
}

var rootCmd = &cobra.Command{
	Use:   "centrifugo-cli",
	Short: "centrifugo-cli is a command line tool for centrifugo",
}

func Exec(args []string) {
	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
		os.Exit(1)
	}
}

// 生成token
func _token() string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix() // jwt过期时间
	claims["iat"] = time.Now().Unix()                                   // jwt签发时间
	token.Claims = claims

	Token, err := token.SignedString([]byte(Secret))
	if err != nil {
		return ""
	}
	return Token
}

func GetHost() string {
	return fmt.Sprintf("%s", Address)
}

func GetPath() string {
	return fmt.Sprintf("/connection/")
}

func Request(method string, path string, data string, header map[string]string) {
	client := http.Client{}
	uri := url.URL{}
	uri.Host = GetHost()
	if Port > 0 && Port < 65536 {
		uri.Host = GetHost() + ":" + strconv.Itoa(Port)
	}
	uri.Scheme = "http"
	uri.Path = path
	logrus.Infof("uri：%s", uri.String())
	req, err := http.NewRequest(method, uri.String(), strings.NewReader(data))
	if err != nil {
		logrus.Error(err)
	}
	req.Header.Set("Authentication", "token "+_token())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "keep-alive")
	for k, v := range header {
		req.Header.Set(k, v)
	}

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

func WebSocket() {

}
