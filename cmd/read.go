package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read message",
	Run: func(cmd *cobra.Command, args []string) {
		go readServer()
	},
}

func readServer() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		Heart(writer)
	})
	logrus.Error(http.ListenAndServe(":10000", nil))
}

func Heart(w http.ResponseWriter) {
	ticker := time.NewTicker(time.Second * time.Duration(3))
	defer func() {
		ticker.Stop()
	}()
	select {
	case msg := <-ReadChan:
		w.Write(msg)
	case <-ticker.C:
		w.Write([]byte(time.Now().Format("2006-01-02 15:04:05") + ": NO DATA"))
	}
}
