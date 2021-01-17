package core

import (
	"testing"

	"github.com/masatrio/parking_lot/vehicle"
)

func Test_parkslot_GetID(t *testing.T) {
	type fields struct {
		id  int
		vhc vehicle.Vehicle
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"get id test", fields{
			id:  1,
			vhc: nil,
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSlot(tt.fields.id)
			if got := s.GetID(); got != tt.want {
				t.Errorf("parkslot.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}
