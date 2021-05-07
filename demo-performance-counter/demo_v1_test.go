package demo_performance_counter

import "testing"

func TestApp_Run(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			"1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := App{}
			a.Run()
		})
	}
}
