package core

import (
	"testing"

	"github.com/masatrio/parking_lot/vehicle/car"

	"github.com/golang/mock/gomock"
	"github.com/masatrio/parking_lot/parkinglot/slot/core"

	"github.com/masatrio/parking_lot/mocks"
	"github.com/masatrio/parking_lot/parkinglot/allocator"
	"github.com/masatrio/parking_lot/parkinglot/slot"
)

func Test_parkingmanager_Leave(t *testing.T) {
	type fields struct {
		allocator   allocator.Allocator
		slotmanager slot.SlotManager
	}
	type args struct {
		id int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	slotManagerMock := mocks.NewMockSlotManager(ctrl)
	allocatorMock := mocks.NewMockAllocator(ctrl)

	const (
		ERRSLOTNOTFOUND = "Error Parking Slot Not Found"
		SUCCESS         = "Success Leave"
	)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{SUCCESS, fields{
			allocator:   allocatorMock,
			slotmanager: slotManagerMock,
		}, args{
			id: 1,
		}, false},
		{ERRSLOTNOTFOUND, fields{
			allocator:   allocatorMock,
			slotmanager: slotManagerMock,
		}, args{
			id: 2,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ParkingManager(allocatorMock, slotManagerMock)

			switch tt.name {
			case SUCCESS:
				{
					slots := make([]slot.Slot, 0)
					slot := core.NewSlot(1)

					vhc := car.Car("123", "blue")

					slot.Fill(vhc)

					slots = append(slots, slot)

					slotManagerMock.EXPECT().GetAll().Return(slots).Times(1)

					slotManagerMock.EXPECT().Dec().Return().Times(1)
				}
			case ERRSLOTNOTFOUND:
				{
					slots := make([]slot.Slot, 0)
					slot := core.NewSlot(1)

					vhc := car.Car("123", "blue")

					slot.Fill(vhc)

					slots = append(slots, slot)

					slotManagerMock.EXPECT().GetAll().Return(slots).Times(1)
				}

			}
			if err := p.Leave(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("parkingmanager.Leave() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
