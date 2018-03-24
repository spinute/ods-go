package meldable_heap

import (
	"github.com/spinute/ods-go/utils"
	"testing"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("MeldableHeap.New().n = %v", ret)
	}
}

func TestAdd(t *testing.T) {
	tests := []utils.V{1, 2, 1, -1, -2, -100, 0, 0}
	mh := New()

	for i, v := range tests {
		mh.Add(v)
		if mh.n != i+1 {
			t.Errorf("mh.n = %v at %v th Add", mh.n, i+1)
		}
	}
}

func TestAddMany(t *testing.T) {
	tests := make([]utils.V, 12345)
	mh := New()

	for i, v := range tests {
		mh.Add(v)
		if mh.n != i+1 {
			t.Errorf("mh.n = %v at %v th Add", mh.n, i+1)
		}
	}
}

func TestAddAndRemove(t *testing.T) {
	tests := []utils.V{1, 2, 1, -1, -2, -100, 0, 0}
	mh := New()

	for _, v := range tests {
		mh.Add(v)
		if ret := mh.Remove(); ret != v {
			t.Errorf("Add %v, then %v was removed", v, ret)
		}
	}
}

func TestAddAndRemove2(t *testing.T) {
	n := 123
	mh := New()

	for i := 0; i < n; i++ {
		mh.Add(utils.V(i))
	}

	for i := 0; i < n; i++ {
		if ret, want := mh.Remove(), utils.V(i); ret != want {
			t.Errorf("expect %d but returned %d", ret, want)
		}
	}
}

func TestAddAndRemove3(t *testing.T) {
	n := 123
	mh := New()

	for i := 0; i < n; i++ {
		mh.Add(utils.V(n - i - 1))
	}

	for i := 0; i < n; i++ {
		if ret, want := mh.Remove(), utils.V(i); ret != want {
			t.Errorf("expect %d but returned %d", ret, want)
		}
	}
}
