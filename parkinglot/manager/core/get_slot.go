package core

import (
	"github.com/masatrio/parking_lot/parkinglot/slot"
)

func (p *parkingmanager) GetSlot(color, regnum string) []slot.Slot {
	slots := p.slotmanager.GetAll()

	newSlots := slots

	result := make([]slot.Slot, 0)

	if regnum != "" {
		for _, val := range newSlots {
			if !val.IsEmpty() {
				if val.GetVehicle().GetRegnum() == regnum {
					result = append(result, val)
					break
				}
			}
		}
	} else if color != "" {
		for _, val := range newSlots {
			if !val.IsEmpty() {
				if val.GetVehicle().GetColor() == color {
					result = append(result, val)
				}
			}
		}
	}

	return result

}
