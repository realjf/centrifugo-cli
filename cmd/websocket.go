package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/url"
	"strconv"
	"time"
)

func WebSocket(cmds []params) {
	uri := url.URL{
		Scheme: "ws",
		Host:   Address,
		Path:   "/connection/websocket",
	}
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

	var dialer *websocket.Dialer
	conn, _, err := dialer.Dial(uri.String(), nil)
	if err != nil {
		logrus.Error(err)
		return
	}

	go timeWriter(conn, &buf)

	for {
		logrus.Info("reading message...")
		_, message, err := conn.ReadMessage()
		if err != nil {
			logrus.Error(err)
			return
		}

		logrus.Infof("received: %s\n", message)
	}
}

func timeWriter(conn *websocket.Conn, buffer *bytes.Buffer) {
	for {
		time.Sleep(time.Second * 2)
		err := conn.WriteMessage(websocket.TextMessage, buffer.Bytes())
		if err != nil {
			logrus.Error(err)
			continue
		}
		break
	}
}
