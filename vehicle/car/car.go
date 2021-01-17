package car

import (
	"github.com/masatrio/parking_lot/vehicle"
)

type car struct {
	color  string
	regnum string
}

func Car(
	regnum string,
	color string,
) vehicle.Vehicle {
	return &car{
		color:  color,
		regnum: regnum,
	}
}
