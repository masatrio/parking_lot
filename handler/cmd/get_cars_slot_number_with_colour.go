package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/masatrio/parking_lot/command"
)

func (h *handle) GetCarsSlotNumberWithColour(cmd command.Command) {
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

	slots := h.parkingmanager.GetSlot(input[1], "")

	if len(slots) == 0 {
		cmd.Error(errors.New("Not found\n"))
		return
	}

	var ids = make([]int, 0)

	for _, slot := range slots {
		ids = append(ids, slot.GetID())
	}

	cmd.Output(fmt.Sprintln(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ", "), "[]")))

	return

}
