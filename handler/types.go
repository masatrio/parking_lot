package handler

import (
	"github.com/masatrio/parking_lot/command"
)

type Handler interface {
	Create(cmd command.Command)
	Park(cmd command.Command)
	Leave(cmd command.Command)
	Status(cmd command.Command)
	GetCarsRegNumberWithColour(cmd command.Command)
	GetCarsSlotNumberWithColour(cmd command.Command)
	GetCarSlotNumberWithRegNum(cmd command.Command)
}

type Validator interface {
	Validate(args []string) error
}
