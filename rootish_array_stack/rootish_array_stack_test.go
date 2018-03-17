package rootish_array_stack

import (
	"github.com/spinute/ods-go/utils"
	"testing"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("RootishArrayStack.New().n = %v", ret)
	}
}

func TestAddHead(t *testing.T) {
	n := 10
	ras := New()

	for i := 0; i < n; i++ {
		ras.Add(0, utils.V(i))
		if ras.n != i+1 {
			t.Errorf("ras.n = %v at %v th Add", ras.n, i+1)
		}
	}
}

func TestAddMid(t *testing.T) {
	n := 10
	ras := New()
	for i := 0; i < n; i++ {
		ras.Add(0, utils.V(i))
	}

	for i := 0; i < n; i++ {
		mid := n / 2
		ras.Add(mid, utils.V(i))
		if ras.n != n+i+1 {
			t.Errorf("ras.n = %v at %v th Add", ras.n, n+i+1)
		}
	}
}

func TestAddLast(t *testing.T) {
	n := 10
	ras := New()

	for i := 0; i < n; i++ {
		ras.Add(i, utils.V(i))
		if ras.n != i+1 {
			t.Errorf("ras.n = %v at %v th Add", ras.n, i+1)
		}
	}
}

func TestGet(t *testing.T) {
	n := 10
	ras := New()
	for i := 0; i < n; i++ {
		ras.Add(i, utils.V(i))
	}

	for i := 0; i < n; i++ {
		if ret := ras.Get(i); ret != utils.V(i) {
			t.Errorf("ras.Get(%v) = %v", i, ret)
		}
	}
}

func TestSet(t *testing.T) {
	n := 10
	ras := New()
	for i := 0; i < n; i++ {
		ras.Add(0, utils.V(i))
	}

	ofs := 1000
	for i := 0; i < n; i++ {
		if ret := ras.Set(i, utils.V(ofs+i)); ret != utils.V(n-i-1) {
			t.Errorf("ras.Set(%v) = %v", i, ret)
		}
	}
	for i := 0; i < n; i++ {
		if ret := ras.Get(i); ret != utils.V(ofs+i) {
			t.Errorf("ras.Get(%v) = %v", i, ret)
		}
	}
}

func TestRemove(t *testing.T) {
	n := 10
	ras := New()
	for i := 0; i < n; i++ {
		ras.Add(0, utils.V(i))
	}

	mid := n / 2
	for i := mid; i < n; i++ {
		if ret := ras.Remove(mid); ret != utils.V(n-i-1) {
			t.Errorf("%v th ras.Remove(mid) = %v", i-mid+1, ret)
		}
	}
	for i := 0; i < mid; i++ {
		if ret := ras.Remove(0); ret != utils.V(n-i-1) {
			t.Errorf("ras.Remove(head) = %v", ret)
		}
	}
}
