package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func WebSocket(cmds []params, header http.Header) interface{} {
	uri := url.URL{
		Scheme: "ws",
		Path:   Path,
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

	logrus.Infof("request uri: %v", uri.String())
	logrus.Infof("request header: %v", header)
	logrus.Infof("request data: %v", cmds)
	var dialer *websocket.Dialer
	conn, _, err := dialer.Dial(uri.String(), header)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	go timeWriter(conn, &buf)

	for {
		logrus.Info("reading message...")
		_, message, err := conn.ReadMessage()
		if err != nil {
			logrus.Error(err)
			return nil
		}

		logrus.Infof("received: %s\n", message)
		return message
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
