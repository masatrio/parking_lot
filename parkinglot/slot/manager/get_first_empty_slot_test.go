package manager

import (
	"reflect"
	"testing"

	"github.com/masatrio/parking_lot/vehicle/car"

	"github.com/masatrio/parking_lot/parkinglot/slot/core"

	"github.com/masatrio/parking_lot/parkinglot/slot"
)

func Test_manager_GetFirstEmptySlot(t *testing.T) {

	successSlots := make([]slot.Slot, 0)
	successSlot := core.NewSlot(1)
	successSlots = append(successSlots, successSlot)

	errorSlots := make([]slot.Slot, 0)
	errorSlot := core.NewSlot(1)
	vhc := car.Car("123", "blue")
	errorSlot.Fill(vhc)
	errorSlots = append(errorSlots, errorSlot)

	type fields struct {
		slots  []slot.Slot
		size   int
		filled int
	}
	tests := []struct {
		name    string
		fields  fields
		want    slot.Slot
		wantErr bool
	}{
		{"get first empty slot success", fields{
			slots:  successSlots,
			filled: 0,
			size:   1,
		}, successSlot, false},
		{"get first empty slot error", fields{
			slots:  errorSlots,
			filled: 1,
			size:   1,
		}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &manager{
				slots:  tt.fields.slots,
				size:   tt.fields.size,
				filled: tt.fields.filled,
			}
			got, err := m.GetFirstEmptySlot()
			if (err != nil) != tt.wantErr {
				t.Errorf("manager.GetFirstEmptySlot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manager.GetFirstEmptySlot() = %v, want %v", got, tt.want)
			}
		})
	}
}
