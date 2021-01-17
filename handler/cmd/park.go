package cmd

import (
	"errors"
	"fmt"

	"github.com/masatrio/parking_lot/vehicle/car"

	"github.com/masatrio/parking_lot/command"
)

func (h *handle) Park(cmd command.Command) {
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

	vhc := car.Car(input[1], input[2])

	slotID, err := h.parkingmanager.Park(vhc)
	if err != nil {
		cmd.Error(err)
		return
	}

	cmd.Output(fmt.Sprintf("Allocated slot number: %v\n", slotID))

	return
}
