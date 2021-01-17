package cmd

import (
	"errors"
	"fmt"

	"github.com/masatrio/parking_lot/command"
)

func (h *handle) GetCarSlotNumberWithRegNum(cmd command.Command) {
	if !h.isParkingInitiated() {
		cmd.Error(errors.New("Please create parking lot first!"))
		return
	}

	input := cmd.GetInput()

	err := h.validator.Validate(input)

	if err != nil {
		cmd.Error(err)
		return
	}

	slots := h.parkingmanager.GetSlot("", input[1])

	if len(slots) == 0 {
		cmd.Error(errors.New("Not found\n"))
		return
	}

	for _, slot := range slots {
		cmd.Output(fmt.Sprintln(slot.GetID()))
	}

	return
}
