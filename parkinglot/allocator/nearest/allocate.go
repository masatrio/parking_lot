package nearest

import (
	"errors"

	"github.com/masatrio/parking_lot/parkinglot/slot"
	"github.com/masatrio/parking_lot/vehicle"
)

func (n *nearest) Allocate(vhc vehicle.Vehicle) (slot.Slot, error) {
	if n.slotManager.IsFull() {
		return nil, errors.New("Cant allocate, slot is full")
	}

	return n.slotManager.GetFirstEmptySlot()
}
