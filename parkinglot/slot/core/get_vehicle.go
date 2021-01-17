package core

import "github.com/masatrio/parking_lot/vehicle"

func (s *parkslot) GetVehicle() vehicle.Vehicle {
	return s.vhc
}
