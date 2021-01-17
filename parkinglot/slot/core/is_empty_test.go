package core

import (
	"testing"

	"github.com/masatrio/parking_lot/vehicle"
)

func Test_parkslot_IsEmpty(t *testing.T) {
	type fields struct {
		id  int
		vhc vehicle.Vehicle
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"is empty test", fields{
			id:  1,
			vhc: nil,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSlot(tt.fields.id)
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("parkslot.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
