package array_queue

import (
	"testing"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("ArrayQueue.New().n = %v", ret)
	}
}

func TestAdd(t *testing.T) {
	tests := [...]Value{1, 2, 1, -1, -2, -100, 0, 0}
	aq := New()

	for i, v := range tests {
		aq.Add(v)
		if aq.n != i+1 {
			t.Errorf("aq.n = %v at %v th Add", aq.n, i+1)
		}
	}
}

func TestAddMany(t *testing.T) {
	tests := make([]Value, 12345)
	aq := New()

	for i, v := range tests {
		aq.Add(v)
		if aq.n != i+1 {
			t.Errorf("aq.n = %v at %v th Add", aq.n, i+1)
		}
	}
}

func TestAddAndRemove(t *testing.T) {
	tests := [...]Value{1, 2, 1, -1, -2, -100, 0, 0}
	aq := New()

	for _, v := range tests {
		aq.Add(v)
		if ret := aq.Remove(); ret != v {
			t.Errorf("Add %v, then %v was removed", v, ret)
		}
	}
}

func TestAddAndRemove2(t *testing.T) {
	n := 123
	aq := New()

	for i := 0; i < n; i++ {
		aq.Add(Value(i))
	}

	for i := 0; i < n; i++ {
		if ret, want := aq.Remove(), Value(i); ret != want {
			t.Errorf("expect %d but returned %d", ret, want)
		}
	}
}
