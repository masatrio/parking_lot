package manager

import (
	"testing"
)

func Test_manager_IsFull(t *testing.T) {
	type fields struct {
		size int
	}

	const (
		FULLTEST    = "full test"
		NOTFULLTEST = "not full test"
	)
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{FULLTEST, fields{
			size: 1,
		}, true},
		{NOTFULLTEST, fields{
			size: 1,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Manager(tt.fields.size)

			switch tt.name {
			case FULLTEST:
				m.Inc()
			case NOTFULLTEST:
			}
			if got := m.IsFull(); got != tt.want {
				t.Errorf("manager.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
