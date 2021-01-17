package manager

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/masatrio/parking_lot/command/core"
)

func (c *cmdMgr) Start(filePath string) {
	var (
		command      string
		err          error
		reader       *bufio.Reader
		readFromFile bool
		args         []string
	)

	if filePath != "" {
		readFromFile = true
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		reader = bufio.NewReader(file)

	} else {

		reader = bufio.NewReader(os.Stdin)

	}

	for {
		if readFromFile {
			line, _, err := reader.ReadLine()

			if err == io.EOF {
				break
			}
			command = fmt.Sprintf("%s", line)

		} else {
			if command, err = reader.ReadString('\n'); err != nil {
				fmt.Println(err)
			}
		}

		args = strings.Split(command, " ")

		for i, v := range args {
			args[i] = strings.TrimSpace(v)
		}

		cmd := core.CMD(args)

		if !c.Handle(cmd) {
			break
		}

		cmd.Print()

	}

}
