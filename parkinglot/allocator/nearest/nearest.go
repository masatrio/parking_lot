package nearest

import (
	"github.com/masatrio/parking_lot/parkinglot/allocator"
	"github.com/masatrio/parking_lot/parkinglot/slot"
)

type nearest struct {
	slotManager slot.SlotManager
}

func Nearest(s slot.SlotManager) allocator.Allocator {
	return &nearest{
		slotManager: s,
	}
}
