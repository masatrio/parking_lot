package cmd

import (
	"github.com/masatrio/parking_lot/handler"
	"github.com/masatrio/parking_lot/parkinglot/manager"
)

type handle struct {
	parkingmanager manager.Manager
	validator      handler.Validator
}

func NewHandle(pm manager.Manager, v handler.Validator) handler.Handler {
	return &handle{
		parkingmanager: pm,
		validator:      v,
	}
}
