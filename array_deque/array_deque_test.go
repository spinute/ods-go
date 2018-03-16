package array_deque

import (
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
		ad.Add(0, Value(i))
		if ad.n != i+1 {
			t.Errorf("ad.n = %v at %v th Add", ad.n, i+1)
		}
	}
}

func TestAddMid(t *testing.T) {
	n := 10
	ad := New()
	for i := 0; i < n; i++ {
		ad.Add(0, Value(i))
	}

	for i := 0; i < n; i++ {
		mid := n / 2
		ad.Add(mid, Value(i))
		if ad.n != n+i+1 {
			t.Errorf("ad.n = %v at %v th Add", ad.n, n+i+1)
		}
	}
}

func TestAddLast(t *testing.T) {
	n := 10
	ad := New()

	for i := 0; i < n; i++ {
		ad.Add(i, Value(i))
		if ad.n != i+1 {
			t.Errorf("ad.n = %v at %v th Add", ad.n, i+1)
		}
	}
}

func TestGet(t *testing.T) {
	n := 10
	ad := New()
	for i := 0; i < n; i++ {
		ad.Add(i, Value(i))
	}

	for i := 0; i < n; i++ {
		if ret := ad.Get(i); ret != Value(i) {
			t.Errorf("ad.Get(%v) = %v", i, ret)
		}
	}
}

func TestSet(t *testing.T) {
	n := 10
	ad := New()
	for i := 0; i < n; i++ {
		ad.Add(0, Value(i))
	}

	ofs := 1000
	for i := 0; i < n; i++ {
		if ret := ad.Set(i, Value(ofs+i)); ret != Value(n-i-1) {
			t.Errorf("ad.Set(%v) = %v", i, ret)
		}
	}
	for i := 0; i < n; i++ {
		if ret := ad.Get(i); ret != Value(ofs+i) {
			t.Errorf("ad.Get(%v) = %v", i, ret)
		}
	}
}

func TestRemove(t *testing.T) {
	n := 10
	ad := New()
	for i := 0; i < n; i++ {
		ad.Add(0, Value(i))
	}

	mid := n / 2
	for i := mid; i < n; i++ {
		if ret := ad.Remove(mid); ret != Value(n-i-1) {
			t.Errorf("%v th ad.Remove(mid) = %v", i-mid+1, ret)
		}
	}
	for i := 0; i < mid; i++ {
		if ret := ad.Remove(0); ret != Value(n-i-1) {
			t.Errorf("ad.Remove(head) = %v", ret)
		}
	}
}
