package allocator

import (
	"github.com/masatrio/parking_lot/parkinglot/slot"
	"github.com/masatrio/parking_lot/vehicle"
)

type Allocator interface {
	Allocate(vhc vehicle.Vehicle) (slot.Slot, error)
}
