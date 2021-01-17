package manager

import (
	"testing"
)

func Test_manager_GetSize(t *testing.T) {
	type fields struct {
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"get size test", fields{
			size: 1,
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Manager(tt.fields.size)
			if got := m.GetSize(); got != tt.want {
				t.Errorf("manager.GetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
