package manager

import (
	"github.com/masatrio/parking_lot/parkinglot/slot"
	"github.com/masatrio/parking_lot/vehicle"
)

type Manager interface {
	Park(vehicle.Vehicle) (int, error)
	Leave(id int) error
	Status() []slot.Slot
	GetSlot(color string, regnum string) []slot.Slot
}
