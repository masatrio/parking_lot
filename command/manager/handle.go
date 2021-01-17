package manager

import (
	"errors"

	"github.com/masatrio/parking_lot/util"

	"github.com/masatrio/parking_lot/command"
)

func (c *cmdMgr) Handle(cmd command.Command) bool {

	input := cmd.GetInput()

	if len(input) < 1 {
		cmd.Error(errors.New(util.Err.NotEnoughArgument))
		return true
	}

	if input[0] == "exit" {
		return false
	}

	m := map[string]func(command.Command){
		"create_parking_lot": c.handle.Create,
		"park":               c.handle.Park,
		"leave":              c.handle.Leave,
		"status":             c.handle.Status,
		"registration_numbers_for_cars_with_colour": c.handle.GetCarsRegNumberWithColour,
		"slot_numbers_for_cars_with_colour":         c.handle.GetCarsSlotNumberWithColour,
		"slot_number_for_registration_number":       c.handle.GetCarSlotNumberWithRegNum,
	}

	if handler, exist := m[input[0]]; exist {
		handler(cmd)
	} else {
		cmd.Error(errors.New("command not identified"))
	}

	return true
}
