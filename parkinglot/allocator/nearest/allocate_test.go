package nearest

import (
	"reflect"
	"testing"

	"github.com/masatrio/parking_lot/parkinglot/slot/core"

	"github.com/masatrio/parking_lot/mocks"
	"github.com/masatrio/parking_lot/vehicle/car"

	"github.com/golang/mock/gomock"
	"github.com/masatrio/parking_lot/parkinglot/slot"
	"github.com/masatrio/parking_lot/vehicle"
)

func Test_nearest_Allocate(t *testing.T) {
	type fields struct {
		slotManager slot.SlotManager
	}
	type args struct {
		vhc vehicle.Vehicle
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	slotManagerMock := mocks.NewMockSlotManager(ctrl)

	vhc := car.Car("123", "blue")

	const (
		TESTERRFULL = "allocate test error full"
		TESTSUCCESS = "allocate test success"
	)

	slotMock := core.NewSlot(1)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    slot.Slot
		wantErr bool
	}{
		{TESTERRFULL, fields{
			slotManager: slotManagerMock,
		}, args{
			vhc: vhc,
		}, nil, true},
		{TESTSUCCESS, fields{
			slotManager: slotManagerMock,
		}, args{
			vhc: vhc,
		}, slotMock, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Nearest(tt.fields.slotManager)

			switch tt.name {
			case TESTERRFULL:
				{
					slotManagerMock.EXPECT().IsFull().Return(true).Times(1)
				}
			case TESTSUCCESS:
				{
					slotManagerMock.EXPECT().IsFull().Return(false).Times(1)
					slotManagerMock.EXPECT().GetFirstEmptySlot().Return(slotMock, nil)
				}
			}

			got, err := n.Allocate(tt.args.vhc)
			if (err != nil) != tt.wantErr {
				t.Errorf("nearest.Allocate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nearest.Allocate() = %v, want %v", got, tt.want)
			}
		})
	}
}
