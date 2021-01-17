package core

import (
	"github.com/masatrio/parking_lot/parkinglot/slot"
	"github.com/masatrio/parking_lot/vehicle"
)

type parkslot struct {
	id  int
	vhc vehicle.Vehicle
}

// NewSlot -
func NewSlot(id int) slot.Slot {
	return &parkslot{
		id: id,
	}
}
