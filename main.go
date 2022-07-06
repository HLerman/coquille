package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/chzyer/readline"
)

func main() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:            "\033[31m»\033[0m ",
		HistoryFile:       ".history",
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
			cmdArgs := strings.Split(cmd, " ")[1:]
			cmd := strings.Split(cmd, " ")[0:1][0]

			if _, err := os.Stat(v + "/" + cmd); err == nil {
				exist = true
				e := exec.Command(v+"/"+cmd, cmdArgs...)
				stdout, err := e.Output()

				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(string(stdout))
				}

				break
			}
		}

		if !exist {
			fmt.Printf("command \"%s\" not found\n", cmd)
		}
	}
}
