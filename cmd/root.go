package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/net/websocket"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&ApiKey, "api-key", "k", "testkey", "set api key for http api")
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
func _token(user string) string {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix() // jwt过期时间
	claims["iat"] = time.Now().Unix()                                   // jwt签发时间
	claims["sub"] = user

	Token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(Secret))
	if err != nil {
		return ""
	}
	return Token
}

func GetHost() string {
	return fmt.Sprintf("%s", Address)
}

type params struct {
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

func Request(method string, path string, cmds []params, header map[string]string) {
	client := http.Client{}
	uri := url.URL{}
	uri.Host = GetHost()
	if Port > 0 && Port < 65536 {
		uri.Host = GetHost() + ":" + strconv.Itoa(Port)
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	for _, cmd := range cmds {
		err := enc.Encode(cmd)
		if err != nil {
			logrus.Error(err)
		}
	}

	uri.Scheme = "http"
	uri.Path = path
	logrus.Infof("uri：%s", uri.String())
	req, err := http.NewRequest(method, uri.String(), &buf)
	if err != nil {
		logrus.Error(err)
	}
	req.Header.Add("Authorization", "apikey "+ApiKey)
	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	logrus.Infof("request header： %v", req.Header)
	resp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Infof("response...")
	logrus.Infof("http status code: %d", resp.StatusCode)
	response, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Infof("%v", string(response))
}

func WebSocket() {
	http.Handle("/upper", websocket.Handler(upper))

	if err := http.ListenAndServe(":9999", nil); err != nil {
		logrus.Info(err)
		os.Exit(1)
	}
}

func upper(ws *websocket.Conn) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			logrus.Info(err)
			continue
		}

		if err = websocket.Message.Send(ws, strings.ToUpper(reply)); err != nil {
			logrus.Info(err)
			continue
		}
	}
}
