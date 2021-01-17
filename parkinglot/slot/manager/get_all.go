package manager

import (
	"github.com/masatrio/parking_lot/parkinglot/slot"
)

func (m *manager) GetAll() []slot.Slot {
	slots := m.slots

	return slots
}
