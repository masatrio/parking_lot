package cmd

import (
	"errors"
	"fmt"

	"github.com/masatrio/parking_lot/command"
)

func (h *handle) Status(cmd command.Command) {
	if !h.isParkingInitiated() {
		cmd.Error(errors.New("Please create parking lot first!"))
		return
	}

	slots := h.parkingmanager.Status()

	statusStr := "Slot No.    Registration No    Colour\n"
	for _, slot := range slots {
		if !slot.IsEmpty() {
			statusStr += fmt.Sprintf("%v           %v      %v\n", slot.GetID(), slot.GetVehicle().GetRegnum(), slot.GetVehicle().GetColor())
		}

	}

	cmd.Output(statusStr)

	return
}
