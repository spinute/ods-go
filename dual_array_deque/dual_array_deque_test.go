package dual_array_deque

import (
	"github.com/spinute/ods-go/utils"
	"testing"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("DualArrayDeque.New().n = %v", ret)
	}
}

func TestAddHead(t *testing.T) {
	n := 10
	dad := New()

	for i := 0; i < n; i++ {
		dad.Add(0, utils.V(i))
		if dad.n != i+1 {
			t.Errorf("dad.n = %v at %v th Add", dad.n, i+1)
		}
	}
}

func TestAddMid(t *testing.T) {
	n := 10
	dad := New()
	for i := 0; i < n; i++ {
		dad.Add(0, utils.V(i))
	}

	for i := 0; i < n; i++ {
		mid := n / 2
		dad.Add(mid, utils.V(i))
		if dad.n != n+i+1 {
			t.Errorf("dad.n = %v at %v th Add", dad.n, n+i+1)
		}
	}
}

func TestAddLast(t *testing.T) {
	n := 10
	dad := New()

	for i := 0; i < n; i++ {
		dad.Add(i, utils.V(i))
		if dad.n != i+1 {
			t.Errorf("dad.n = %v at %v th Add", dad.n, i+1)
		}
	}
}

func TestGet(t *testing.T) {
	n := 10
	dad := New()
	for i := 0; i < n; i++ {
		dad.Add(i, utils.V(i))
	}

	for i := 0; i < n; i++ {
		if ret := dad.Get(i); ret != utils.V(i) {
			t.Errorf("dad.Get(%v) = %v", i, ret)
		}
	}
}

func TestSet(t *testing.T) {
	n := 10
	dad := New()
	for i := 0; i < n; i++ {
		dad.Add(0, utils.V(i))
	}

	ofs := 1000
	for i := 0; i < n; i++ {
		if ret := dad.Set(i, utils.V(ofs+i)); ret != utils.V(n-i-1) {
			t.Errorf("dad.Set(%v) = %v", i, ret)
		}
	}
	for i := 0; i < n; i++ {
		if ret := dad.Get(i); ret != utils.V(ofs+i) {
			t.Errorf("dad.Get(%v) = %v", i, ret)
		}
	}
}

func TestRemove(t *testing.T) {
	n := 10
	dad := New()
	for i := 0; i < n; i++ {
		dad.Add(0, utils.V(i))
	}

	mid := n / 2
	for i := mid; i < n; i++ {
		if ret := dad.Remove(mid); ret != utils.V(n-i-1) {
			t.Errorf("%v th dad.Remove(mid) = %v", i-mid+1, ret)
		}
	}
	for i := 0; i < mid; i++ {
		if ret := dad.Remove(0); ret != utils.V(n-i-1) {
			t.Errorf("dad.Remove(head) = %v", ret)
		}
	}
}
