package core

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/masatrio/parking_lot/mocks"
	"github.com/masatrio/parking_lot/parkinglot/allocator"
	"github.com/masatrio/parking_lot/parkinglot/slot"
)

func Test_parkingmanager_Status(t *testing.T) {
	type fields struct {
		allocator   allocator.Allocator
		slotmanager slot.SlotManager
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	slotManagerMock := mocks.NewMockSlotManager(ctrl)
	allocatorMock := mocks.NewMockAllocator(ctrl)

	retSuccess := make([]slot.Slot, 0)
	tests := []struct {
		name   string
		fields fields
		want   []slot.Slot
	}{
		{"Success", fields{
			allocator:   allocatorMock,
			slotmanager: slotManagerMock,
		}, retSuccess},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ParkingManager(tt.fields.allocator, tt.fields.slotmanager)

			slotManagerMock.EXPECT().GetAll().Return(retSuccess).Times(1)

			if got := p.Status(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parkingmanager.Status() = %v, want %v", got, tt.want)
			}
		})
	}
}
