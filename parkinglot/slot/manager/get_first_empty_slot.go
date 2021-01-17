package manager

import (
	"errors"

	"github.com/masatrio/parking_lot/parkinglot/slot"
)

func (m *manager) GetFirstEmptySlot() (slot.Slot, error) {

	for _, slot := range m.slots {
		if slot.IsEmpty() {
			return slot, nil
		}
	}

	return nil, errors.New("Not Found")
}
