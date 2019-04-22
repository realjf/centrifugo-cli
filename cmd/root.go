package cmd

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"time"
)

var Secret string
var ApiKey string
var Token string

func init() {
	rootCmd.PersistentFlags().StringVarP(&ApiKey, "api-key", "k", "", "set api key for http api")
	rootCmd.PersistentFlags().StringVarP(&Secret, "secret", "s", "", "secret key for http api")
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
