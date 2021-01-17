package manager

import (
	"testing"
)

func Test_manager_Dec(t *testing.T) {
	type fields struct {
		size int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"decrement test", fields{
			size: 2,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Manager(tt.fields.size)
			m.Inc()
			m.Dec()
		})
	}
}
