package core

import (
	"testing"

	"github.com/masatrio/parking_lot/vehicle/car"

	"github.com/masatrio/parking_lot/vehicle"
)

func Test_parkslot_Fill(t *testing.T) {
	type fields struct {
		id  int
		vhc vehicle.Vehicle
	}
	type args struct {
		vhc vehicle.Vehicle
	}

	vhc := car.Car("123", "blue")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"fill test success", fields{
			id:  1,
			vhc: vhc,
		}, args{
			vhc: vhc,
		}, false},
		{"fill test error", fields{
			id:  1,
			vhc: nil,
		}, args{
			vhc: nil,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSlot(tt.fields.id)
			if err := s.Fill(tt.args.vhc); (err != nil) != tt.wantErr {
				t.Errorf("parkslot.Fill() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
