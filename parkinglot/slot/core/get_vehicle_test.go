package core

import (
	"reflect"
	"testing"

	"github.com/masatrio/parking_lot/vehicle/car"

	"github.com/masatrio/parking_lot/vehicle"
)

func Test_parkslot_GetVehicle(t *testing.T) {
	type fields struct {
		id  int
		vhc vehicle.Vehicle
	}

	vhc := car.Car("123", "blue")

	tests := []struct {
		name   string
		fields fields
		want   vehicle.Vehicle
	}{
		{"get vehicle test", fields{
			id:  1,
			vhc: vhc,
		}, vhc},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSlot(tt.fields.id)
			s.Fill(vhc)
			if got := s.GetVehicle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parkslot.GetVehicle() = %v, want %v", got, tt.want)
			}
		})
	}
}
