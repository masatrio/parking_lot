package core

import (
	"errors"
	"testing"

	"github.com/masatrio/parking_lot/parkinglot/slot/core"

	"github.com/golang/mock/gomock"
	"github.com/masatrio/parking_lot/mocks"
	"github.com/masatrio/parking_lot/parkinglot/allocator"
	"github.com/masatrio/parking_lot/parkinglot/slot"
	"github.com/masatrio/parking_lot/vehicle"
	"github.com/masatrio/parking_lot/vehicle/car"
)

func Test_parkingmanager_Park(t *testing.T) {
	type fields struct {
		allocator   allocator.Allocator
		slotmanager slot.SlotManager
	}

	type args struct {
		vhc vehicle.Vehicle
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	slotManagerMock := mocks.NewMockSlotManager(ctrl)
	allocatorMock := mocks.NewMockAllocator(ctrl)

	vhc := car.Car("123", "blue")

	const (
		ERRPARKFULL = "error parking lot full"
		SUCCESSPARK = "success park"
	)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{ERRPARKFULL, fields{
			slotmanager: slotManagerMock,
			allocator:   allocatorMock,
		}, args{
			vhc: vhc,
		}, 0, true},
		{SUCCESSPARK, fields{
			slotmanager: slotManagerMock,
			allocator:   allocatorMock,
		}, args{
			vhc: vhc,
		}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ParkingManager(tt.fields.allocator, tt.fields.slotmanager)

			switch tt.name {
			case ERRPARKFULL:
				allocatorMock.EXPECT().Allocate(tt.args.vhc).Return(nil, errors.New("Cant allocate, slot is full")).Times(1)
			case SUCCESSPARK:
				{
					newSlot := core.NewSlot(1)
					allocatorMock.EXPECT().Allocate(tt.args.vhc).Return(newSlot, nil).Times(1)
					slotManagerMock.EXPECT().Inc().Return().Times(1)
				}
			}

			got, err := p.Park(tt.args.vhc)
			if (err != nil) != tt.wantErr {
				t.Errorf("parkingmanager.Park() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parkingmanager.Park() = %v, want %v", got, tt.want)
			}
		})
	}
}
