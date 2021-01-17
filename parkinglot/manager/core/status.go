package core

import "github.com/masatrio/parking_lot/parkinglot/slot"

func (p *parkingmanager) Status() []slot.Slot {

	slots := p.slotmanager.GetAll()
	newSlot := slots
	return newSlot
}
