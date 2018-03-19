package skiplistlist

import (
	"github.com/spinute/ods-go/utils"
	"testing"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("SkiplistList.New().n = %v", ret)
	}
}

func TestAddHead(t *testing.T) {
	n := 10
	sll := New()

	for i := 0; i < n; i++ {
		sll.Add(0, utils.V(i))
		if sll.n != i+1 {
			t.Errorf("sll.n = %v at %v th Add", sll.n, i+1)
		}
	}
}

func TestAddMid(t *testing.T) {
	n := 10
	sll := New()
	for i := 0; i < n; i++ {
		sll.Add(0, utils.V(i))
	}

	for i := 0; i < n; i++ {
		mid := n / 2
		sll.Add(mid, utils.V(i))
		if sll.n != n+i+1 {
			t.Errorf("sll.n = %v at %v th Add", sll.n, n+i+1)
		}
	}
}

func TestAddLast(t *testing.T) {
	n := 10
	sll := New()

	for i := 0; i < n; i++ {
		sll.Add(i, utils.V(i))
		if sll.n != i+1 {
			t.Errorf("sll.n = %v at %v th Add", sll.n, i+1)
		}
	}
}

func TestGet(t *testing.T) {
	n := 10
	sll := New()
	for i := 0; i < n; i++ {
		sll.Add(i, utils.V(i))
	}

	for i := 0; i < n; i++ {
		if ret := sll.Get(i); ret != utils.V(i) {
			t.Errorf("sll.Get(%v) = %v", i, ret)
		}
	}
}

func TestSet(t *testing.T) {
	n := 10
	sll := New()
	for i := 0; i < n; i++ {
		sll.Add(0, utils.V(i))
	}

	ofs := 1000
	for i := 0; i < n; i++ {
		if ret := sll.Set(i, utils.V(ofs+i)); ret != utils.V(n-i-1) {
			t.Errorf("sll.Set(%v) = %v", i, ret)
		}
	}
	for i := 0; i < n; i++ {
		if ret := sll.Get(i); ret != utils.V(ofs+i) {
			t.Errorf("sll.Get(%v) = %v", i, ret)
		}
	}
}

func TestRemove(t *testing.T) {
	n := 10
	sll := New()
	for i := 0; i < n; i++ {
		sll.Add(0, utils.V(i))
	}

	mid := n / 2
	for i := mid; i < n; i++ {
		if ret := sll.Remove(mid); ret != utils.V(n-i-1) {
			t.Errorf("%v th sll.Remove(mid) = %v", i-mid+1, ret)
		}
	}
	for i := 0; i < mid; i++ {
		if ret := sll.Remove(0); ret != utils.V(n-i-1) {
			t.Errorf("sll.Remove(head) = %v", ret)
		}
	}
}
