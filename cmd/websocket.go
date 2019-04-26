package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Hour

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512 * 100000
)

func WebSocket(cmds []params, header http.Header) {
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
	CreateConn()

	// 检查websocket链接
	go timeWriter(WebSocketConn, &buf)
	go ReadPipe(WebSocketConn)
}

func CreateConn() {
	var err error
	if WebSocketConn == nil {
		WebSocketConn, err = websocket.Dial(URI.String(), "", URI.Host)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func ReadPipe(conn *websocket.Conn) {
	conn.SetReadDeadline(time.Now().Add(pongWait))
	for {
		var message []byte
		err := websocket.Message.Receive(conn, &message)
		if err != nil {
			if err == io.EOF {
				time.Sleep(time.Second * 3)
				continue
			}
			logrus.Error(err)
			break
		}
		logrus.Info("reading message...")
		logrus.Infof("received: %s\n", message)
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
			logrus.Error(err)
		}
		if data.Result.Client != "" {
			ClientConnectionID = data.Result.Client
		}

		ReadChan <- message
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
	conn.SetWriteDeadline(time.Now().Add(pongWait))
	for {
		time.Sleep(time.Second * 2)
		if buffer.Len() > 0 {
			err := websocket.Message.Send(conn, buffer.Bytes())
			if err != nil {
				logrus.Error(err)
				WebSocketConn, err = websocket.Dial(URI.String(), "", URI.Host)
				if err != nil {
					logrus.Error(err)
				}
			}
			buffer.Truncate(0)
		}
	}
}
