package core

import (
	"github.com/masatrio/parking_lot/parkinglot/allocator"
	"github.com/masatrio/parking_lot/parkinglot/manager"
	"github.com/masatrio/parking_lot/parkinglot/slot"
)

type parkingmanager struct {
	allocator   allocator.Allocator
	slotmanager slot.SlotManager
}

func ParkingManager(
	allocator allocator.Allocator,
	slotmanager slot.SlotManager,
) manager.Manager {
	return &parkingmanager{
		allocator:   allocator,
		slotmanager: slotmanager,
	}
}
