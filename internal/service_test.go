package internal

import "testing"

func Test_controller(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "unit test"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s pass", tt.name)
		})
	}
}
