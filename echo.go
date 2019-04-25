package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	handle := http.Handle()
	logrus.Error(http.ListenAndServe(":10000", handle))
}
