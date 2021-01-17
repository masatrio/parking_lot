package manager

import (
	"reflect"
	"testing"

	"github.com/masatrio/parking_lot/parkinglot/slot"
)

func Test_manager_GetAll(t *testing.T) {
	type fields struct {
		slots  []slot.Slot
		size   int
		filled int
	}
	tests := []struct {
		name   string
		fields fields
		want   []slot.Slot
	}{
		{"get all slot", fields{
			size:   1,
			slots:  make([]slot.Slot, 0),
			filled: 0,
		}, make([]slot.Slot, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &manager{
				slots:  tt.fields.slots,
				size:   tt.fields.size,
				filled: tt.fields.filled,
			}
			if got := m.GetAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manager.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
