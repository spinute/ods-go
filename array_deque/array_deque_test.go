package array_deque

import (
	"github.com/spinute/ods-go/utils"
	"testing"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("ArrayDeque.New().n = %v", ret)
	}
}

func TestAddHead(t *testing.T) {
	n := 10
	ad := New()

	for i := 0; i < n; i++ {
		ad.Add(0, utils.V(i))
		if ad.n != i+1 {
			t.Errorf("ad.n = %v at %v th Add", ad.n, i+1)
		}
	}
}

func TestAddMid(t *testing.T) {
	n := 10
	ad := New()
	for i := 0; i < n; i++ {
		ad.Add(0, utils.V(i))
	}

	for i := 0; i < n; i++ {
		mid := n / 2
		ad.Add(mid, utils.V(i))
		if ad.n != n+i+1 {
			t.Errorf("ad.n = %v at %v th Add", ad.n, n+i+1)
		}
	}
}

func TestAddLast(t *testing.T) {
	n := 10
	ad := New()

	for i := 0; i < n; i++ {
		ad.Add(i, utils.V(i))
		if ad.n != i+1 {
			t.Errorf("ad.n = %v at %v th Add", ad.n, i+1)
		}
	}
}

func TestGet(t *testing.T) {
	n := 10
	ad := New()
	for i := 0; i < n; i++ {
		ad.Add(i, utils.V(i))
	}

	for i := 0; i < n; i++ {
		if ret := ad.Get(i); ret != utils.V(i) {
			t.Errorf("ad.Get(%v) = %v", i, ret)
		}
	}
}

func TestSet(t *testing.T) {
	n := 10
	ad := New()
	for i := 0; i < n; i++ {
		ad.Add(0, utils.V(i))
	}

	ofs := 1000
	for i := 0; i < n; i++ {
		if ret := ad.Set(i, utils.V(ofs+i)); ret != utils.V(n-i-1) {
			t.Errorf("ad.Set(%v) = %v", i, ret)
		}
	}
	for i := 0; i < n; i++ {
		if ret := ad.Get(i); ret != utils.V(ofs+i) {
			t.Errorf("ad.Get(%v) = %v", i, ret)
		}
	}
}

func TestRemove(t *testing.T) {
	n := 10
	ad := New()
	for i := 0; i < n; i++ {
		ad.Add(0, utils.V(i))
	}

	mid := n / 2
	for i := mid; i < n; i++ {
		if ret := ad.Remove(mid); ret != utils.V(n-i-1) {
			t.Errorf("%v th ad.Remove(mid) = %v", i-mid+1, ret)
		}
	}
	for i := 0; i < mid; i++ {
		if ret := ad.Remove(0); ret != utils.V(n-i-1) {
			t.Errorf("ad.Remove(head) = %v", ret)
		}
	}
}
