package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:            "\033[31m»\033[0m ",
		HistoryFile:       "readline.tmp",
		InterruptPrompt:   "^C",
		EOFPrompt:         "exit",
		HistorySearchFold: true,
	})

	if err != nil {
		panic(err)
	}

	defer l.Close()

	for {
		cmd, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(cmd) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		p := strings.Split(os.Getenv("PATH"), ":")
		exist := false
		for _, v := range p {
			if _, err := os.Stat(v + "/" + cmd); err == nil {
				exist = true
				fmt.Println(v + "/" + cmd)
			}
		}

		if !exist {
			fmt.Printf("command \"%s\" not found\n", cmd)
		}
	}
}
