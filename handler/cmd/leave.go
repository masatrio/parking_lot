package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/masatrio/parking_lot/util"

	"github.com/masatrio/parking_lot/command"
)

func (h *handle) Leave(cmd command.Command) {
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

	slotNumber, err := strconv.Atoi(input[1])
	if err != nil {
		cmd.Error(errors.New(util.Err.WrongInputFormat))
		return
	}

	err = h.parkingmanager.Leave(slotNumber)
	if err != nil {
		cmd.Error(err)
		return
	}

	cmd.Output(fmt.Sprintf("Slot number %v is free\n", slotNumber))

	return
}
