package core

import (
	"reflect"
	"testing"

	"github.com/masatrio/parking_lot/parkinglot/slot/core"
	"github.com/masatrio/parking_lot/vehicle/car"

	"github.com/golang/mock/gomock"
	"github.com/masatrio/parking_lot/mocks"
	"github.com/masatrio/parking_lot/parkinglot/allocator"
	"github.com/masatrio/parking_lot/parkinglot/slot"
)

func Test_parkingmanager_GetSlot(t *testing.T) {
	type fields struct {
		allocator   allocator.Allocator
		slotmanager slot.SlotManager
	}
	type args struct {
		color  string
		regnum string
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	slotManagerMock := mocks.NewMockSlotManager(ctrl)
	allocatorMock := mocks.NewMockAllocator(ctrl)

	const (
		GETBYREGNUM        = "Get By Regnum"
		GETBYCOLOR         = "Get By Color"
		GETBYCOLORNOTFOUND = "Get By Color Not Found"
	)

	vhc := car.Car("123", "blue")
	newSlot := core.NewSlot(1)
	newSlot.Fill(vhc)
	ret := make([]slot.Slot, 0)
	ret = append(ret, newSlot)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   []slot.Slot
	}{
		{GETBYREGNUM, fields{
			allocator:   allocatorMock,
			slotmanager: slotManagerMock,
		}, args{
			color:  "",
			regnum: "123",
		}, ret},
		{GETBYCOLOR, fields{
			allocator:   allocatorMock,
			slotmanager: slotManagerMock,
		}, args{
			color:  "blue",
			regnum: "",
		}, ret},
		{GETBYCOLORNOTFOUND, fields{
			allocator:   allocatorMock,
			slotmanager: slotManagerMock,
		}, args{
			color:  "yellow",
			regnum: "",
		}, make([]slot.Slot, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ParkingManager(allocatorMock, slotManagerMock)

			switch tt.name {
			case GETBYREGNUM:
				{
					slotManagerMock.EXPECT().GetAll().Return(ret).Times(1)
				}
			case GETBYCOLOR:
				{
					slotManagerMock.EXPECT().GetAll().Return(ret).Times(1)
				}
			case GETBYCOLORNOTFOUND:
				{
					slotManagerMock.EXPECT().GetAll().Return(ret).Times(1)
				}
			}
			if got := p.GetSlot(tt.args.color, tt.args.regnum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parkingmanager.GetSlot() = %v, want %v", got, tt.want)
			}
		})
	}
}
