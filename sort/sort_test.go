package sort

import (
	"testing"

	"github.com/spinute/ods-go/utils"
)

func slice_equal(a, b []utils.V) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		arg      []utils.V
		expected []utils.V
	}{
		{[]utils.V{}, []utils.V{}},
		{[]utils.V{1}, []utils.V{1}},
		{[]utils.V{0, 1}, []utils.V{0, 1}},
		{[]utils.V{1, 0, 1}, []utils.V{0, 1, 1}},
		{[]utils.V{3, 2, -1}, []utils.V{-1, 2, 3}},
		{[]utils.V{3, 2, -1, 3, 2, 1, 1, 2, 3}, []utils.V{-1, 1, 1, 2, 2, 2, 3, 3, 3}},
	}
	for _, test := range tests {
		if ret := MergeSort(test.arg); slice_equal(ret, test.expected) == false {
			t.Errorf("MergeSort(%v) = %v", test.arg, ret)
		}
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		arg      []utils.V
		expected []utils.V
	}{
		{[]utils.V{}, []utils.V{}},
		{[]utils.V{1}, []utils.V{1}},
		{[]utils.V{0, 1}, []utils.V{0, 1}},
		{[]utils.V{1, 0, 1}, []utils.V{0, 1, 1}},
		{[]utils.V{3, 2, -1}, []utils.V{-1, 2, 3}},
		{[]utils.V{3, 2, -1, 3, 2, 1, 1, 2, 3}, []utils.V{-1, 1, 1, 2, 2, 2, 3, 3, 3}},
	}
	for _, test := range tests {
		if ret := QuickSort(test.arg); slice_equal(ret, test.expected) == false {
			t.Errorf("QuickSort(%v) = %v", test.arg, ret)
		}
	}
}
