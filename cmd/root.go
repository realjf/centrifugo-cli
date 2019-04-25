package cmd

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		logrus.Error(err)
	}
}

// 生成token
func _token(user uint32) string {
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
	ID     uint32                 `json:"id"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}
