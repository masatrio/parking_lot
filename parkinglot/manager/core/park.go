package core

import (
	"errors"
	"fmt"

	"github.com/masatrio/parking_lot/vehicle"
)

func (p *parkingmanager) Park(vhc vehicle.Vehicle) (int, error) {
	slot, err := p.allocator.Allocate(vhc)
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Sorry, parking lot is full\n"))
	}

	err = slot.Fill(vhc)
	if err == nil {
		p.slotmanager.Inc()
	}

	return slot.GetID(), err

}
