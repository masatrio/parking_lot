package manager

import (
	"github.com/masatrio/parking_lot/parkinglot/slot"
	"github.com/masatrio/parking_lot/parkinglot/slot/core"
)

type manager struct {
	slots  []slot.Slot
	size   int
	filled int
}

func Manager(size int) slot.SlotManager {

	var slots = make([]slot.Slot, size)

	for i := 0; i < size; i++ {
		slots[i] = core.NewSlot(i + 1)
	}

	return &manager{
		size:   size,
		slots:  slots,
		filled: 0,
	}
}
