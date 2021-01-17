package core

import (
	"errors"
)

func (s *parkslot) Empty() error {
	if s.vhc == nil {
		return errors.New("Slot already empty")
	}

	s.vhc = nil

	return nil
}
