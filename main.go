package main

import (
	"bufio"
	"centrifugo-cli/cmd"
	"fmt"
	"os"
	"strings"
)

func init() {

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	cmd.ReadChan = make(chan []byte, 204800)
	for {
		fmt.Print("$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = Run(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func Run(cmdString string) error {
	cmdString = strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(cmdString)
	if len(arrCommandStr) > 0 {
		switch arrCommandStr[0] {
		case "exit":
			os.Exit(0)
		}
	}

	cmd.Exec(arrCommandStr)

	return nil
}
