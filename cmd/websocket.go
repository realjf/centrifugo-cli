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

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512 * 100000
)

func WebSocket(cmds []params, header http.Header) interface{} {
	URI = url.URL{
		Scheme: "ws",
		Path:   Path,
	}
	URI.Host = GetHost()
	if Port > 0 && Port < 65536 {
		URI.Host = GetHost() + ":" + strconv.Itoa(Port)
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	for _, cmd := range cmds {
		err := enc.Encode(cmd)
		if err != nil {
			logrus.Error(err)
		}
	}

	logrus.Infof("request uri: %v", URI.String())
	logrus.Infof("request header: %v", header)
	logrus.Infof("request data: %v", cmds)
	CreateConn(URI, header)

	// 检查websocket链接
	go HeartCheck(WebSocketConn)
	go timeWriter(WebSocketConn, &buf)
	return ReadPipe(WebSocketConn)
}

func CreateConn(uri url.URL, header http.Header) {
	var err error
	if WebSocketConn == nil {
		var dialer *websocket.Dialer
		WebSocketConn, _, err = dialer.Dial(uri.String(), header)
		if err != nil {
			logrus.Error(err)
		}
	}
	err = WebSocketConn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		logrus.Error(err)
	}
}

func ReadPipe(conn *websocket.Conn) interface{} {
	for {
		logrus.Info("reading message...")

		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logrus.Error(err)
				return nil
			}
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				logrus.Error(err)
				return nil
			}
			logrus.Error(err)
		}

		logrus.Infof("received: %s\n", message)
		return message
	}
}

// 心跳检测
func HeartCheck(conn *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case <-ticker.C:
			if err := conn.SetReadDeadline(time.Now().Add(writeWait)); err != nil {
				logrus.Error(err)
				break
			}
		}
	}
}

func timeWriter(conn *websocket.Conn, buffer *bytes.Buffer) {
	for {
		time.Sleep(time.Second * 2)
		err := conn.WriteMessage(websocket.TextMessage, buffer.Bytes())
		if err != nil {
			logrus.Error(err)
			return
		}
		break
	}
}
