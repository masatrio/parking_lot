package core

import (
	"fmt"

	"github.com/masatrio/parking_lot/command"
)

type cmd struct {
	input  []string
	output string
	err    error
}

func CMD(input []string) command.Command {
	return &cmd{
		input: input,
	}

}

func (c *cmd) GetInput() []string {
	return c.input
}

func (c *cmd) GetOutput() string {
	return c.output
}

func (c *cmd) GetErr() error {
	return c.err
}

func (c *cmd) Error(err error) {
	c.err = err
}

func (c *cmd) Output(str string) {
	c.output = str
}

func (c *cmd) Print() {
	str := ""
	if c.GetErr() != nil {
		str = c.GetErr().Error()
	} else {
		str = c.GetOutput()
	}

	fmt.Print(str)
}
