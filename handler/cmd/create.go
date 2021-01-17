package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/masatrio/parking_lot/util"

	"github.com/masatrio/parking_lot/parkinglot/allocator/nearest"
	"github.com/masatrio/parking_lot/parkinglot/manager/core"
	"github.com/masatrio/parking_lot/parkinglot/slot/manager"

	"github.com/masatrio/parking_lot/command"
)

func (h *handle) Create(cmd command.Command) {

	if h.isParkingInitiated() {
		cmd.Error(errors.New("Parking lot already created !\n"))
		return
	}

	input := cmd.GetInput()

	err := h.validator.Validate(input)

	if err != nil {
		cmd.Error(err)
		return
	}

	size, err := strconv.Atoi(input[1])
	if err != nil {
		cmd.Error(errors.New(util.Err.WrongInputFormat))
		return
	}

	if size < 1 {
		cmd.Error(errors.New("Parking size must be more than zero !\n"))
		return
	}

	slotMgr := manager.Manager(size)

	allocator := nearest.Nearest(slotMgr)

	parkingmanager := core.ParkingManager(allocator, slotMgr)

	h.parkingmanager = parkingmanager

	cmd.Output(fmt.Sprintf("Created a parking lot with %v slots\n", size))

	return
}
