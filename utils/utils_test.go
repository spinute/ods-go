package utils

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		arg1, arg2, want int
	}{
		{0, 1, 1},
		{-123, 321, 321},
		{3, 3, 3},
	}

	for _, test := range tests {
		if ret := Max(test.arg1, test.arg2); ret != test.want {
			t.Errorf("Max(%v, %v) = %v", test.arg1, test.arg2, ret)
		}
	}
}
