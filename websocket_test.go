package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/websocket"
	//"net/url"
	"testing"
	"time"
)

var Url = "ws://localhost:8000/connection/websocket"
var Origin = "http://localhost:8000/"

func TestWebSock(t *testing.T) {
	websock()
	t.Error("conn")
}

func websock() {
	//uri := url.URL{
	//	Scheme: "ws",
	//	Host:   "localhost:8000",
	//	Path:   "/connection/websocket",
	//}

	ws, err := websocket.Dial(Url, "", Origin)
	if err != nil {
		fmt.Println(err)
		return
	}

	go timeWriter(ws)

	for {
		var message string
		err := websocket.Message.Receive(ws, &message)
		if err != nil {
			fmt.Println("read:", err)
			return
		}

		fmt.Printf("received: %s\n", message)
		var data struct {
			ID     int `json:"id"`
			Result struct {
				Client  string `json:"client"`
				Version string `json:"version"`
				Expires bool   `json:"expires"`
				Ttl     int    `json:"ttl"`
			}
		}
		err = json.Unmarshal([]byte(message), &data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("received: %v\n", data)
	}
}

func timeWriter(ws *websocket.Conn) {
	for {
		time.Sleep(time.Second * 2)
		hell := map[string]interface{}{
			"id":     1,
			"method": "connect",
			"params": map[string]interface{}{
				"token": _token("1"),
				"data": struct {
				}{},
			},
		}
		data, _ := json.Marshal(&hell)
		err := websocket.Message.Send(ws, data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// 生成token
func _token(user string) string {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix() // jwt过期时间
	claims["iat"] = time.Now().Unix()                                   // jwt签发时间
	claims["sub"] = user

	Token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("178e1a29-fbc8-4422-a47e-3b66a411677b"))
	if err != nil {
		return ""
	}
	return Token
}
