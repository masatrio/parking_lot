package slot

import (
	"github.com/masatrio/parking_lot/vehicle"
)

type Slot interface {
	GetID() int
	GetVehicle() vehicle.Vehicle
	Fill(vehicle.Vehicle) error
	Empty() error
	IsEmpty() bool
}

type SlotManager interface {
	GetSize() int
	GetAll() []Slot
	GetFirstEmptySlot() (Slot, error)
	Inc()
	Dec()
	IsFull() bool
}
