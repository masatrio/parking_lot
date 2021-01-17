package main

import (
	"log"
	"os"

	"github.com/masatrio/parking_lot/command"
	"github.com/masatrio/parking_lot/command/manager"

	"github.com/masatrio/parking_lot/handler/cmd"
)

var (
	cmdMgr   command.CommandManager
	filePath string
)

func init() {

	// get os args
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	// init handler
	validator := cmd.Validator()
	handler := cmd.NewHandle(nil, validator)

	// init command manager
	cmdMgr = manager.CMDMgr(handler)
}

func main() {
	if err := cmdMgr.Start(filePath); err != nil {
		log.Fatal(err)
	}
}
