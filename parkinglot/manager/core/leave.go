package core

import (
	"errors"

	"github.com/masatrio/parking_lot/parkinglot/slot"
)

func (p *parkingmanager) Leave(id int) error {
	slots := p.slotmanager.GetAll()
	var target slot.Slot

	for _, slot := range slots {
		if slot.GetID() == id {
			target = slot
			break
		}
	}

	if target == nil {
		return errors.New("Slot Not Found\n")
	}

	err := target.Empty()
	if err == nil {
		p.slotmanager.Dec()
	}

	return err

}
