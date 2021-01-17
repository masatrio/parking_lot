package core

import (
	"errors"

	"github.com/masatrio/parking_lot/vehicle"
)

func (s *parkslot) Fill(vhc vehicle.Vehicle) error {
	if vhc == nil {
		return errors.New("Vehicle cannot be nil")
	}

	s.vhc = vhc

	return nil
}
