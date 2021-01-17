package core

import (
	"testing"

	"github.com/masatrio/parking_lot/vehicle/car"

	"github.com/masatrio/parking_lot/vehicle"
)

func Test_parkslot_Empty(t *testing.T) {
	type fields struct {
		id  int
		vhc vehicle.Vehicle
	}

	vhc := car.Car("123", "blue")

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"empty test success", fields{
			id:  1,
			vhc: vhc,
		}, false},
		{"empty test error", fields{
			id:  1,
			vhc: nil,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSlot(tt.fields.id)
			s.Fill(tt.fields.vhc)
			if err := s.Empty(); (err != nil) != tt.wantErr {
				t.Errorf("parkslot.Empty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
