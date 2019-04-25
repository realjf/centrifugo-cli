package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"net/url"
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
	uri := url.URL{
		Scheme: "ws",
		Host:   "localhost:8000",
		Path:   "/connection/websocket",
	}

	var dialer *websocket.Dialer
	conn, _, err := dialer.Dial(uri.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go timeWriter(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}

		fmt.Printf("received: %s\n", message)
	}
}

func timeWriter(conn *websocket.Conn) {
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
		conn.WriteMessage(websocket.TextMessage, []byte(data))
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
