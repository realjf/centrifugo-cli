package cmd

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

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
