package manager

import (
	"github.com/masatrio/parking_lot/command"
	"github.com/masatrio/parking_lot/handler"
)

type cmdMgr struct {
	handle handler.Handler
}

func CMDMgr(h handler.Handler) command.CommandManager {
	return &cmdMgr{
		handle: h,
	}
}
