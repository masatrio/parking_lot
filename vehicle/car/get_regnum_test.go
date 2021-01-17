package car

import "testing"

func Test_car_GetRegnum(t *testing.T) {
	type fields struct {
		color  string
		regnum string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test car get regnum", fields{
			color:  "blue",
			regnum: "123",
		}, "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Car(tt.fields.regnum, tt.fields.color)
			if got := c.GetRegnum(); got != tt.want {
				t.Errorf("car.GetRegnum() = %v, want %v", got, tt.want)
			}
		})
	}
}
