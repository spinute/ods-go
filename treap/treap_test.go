package treap

import (
	"testing"

	"github.com/spinute/ods-go/utils"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("Treap.New().n = %v", ret)
	}
}

func TestAdd(t *testing.T) {
	n := 10
	treap := New()

	for i := 0; i < n; i++ {
		ret := treap.Add(utils.V(i))
		if treap.n != i+1 {
			t.Errorf("treap.n = %v at %v th Add", treap.n, i+1)
		}
		if ret == false {
			t.Errorf("Add returned false unexpectedly")
		}
	}

	for i := 0; i < n; i++ {
		ret := treap.Add(utils.V(i))
		if treap.n != n {
			t.Errorf("treap.n = %v", treap.n)
		}
		if ret == true {
			t.Errorf("Add returned true unexpectedly")
		}
	}
}

func TestFindEQ(t *testing.T) {
	n := 10
	treap := New()

	for i := 0; i < n; i++ {
		ret := treap.FindEQ(utils.V(i))
		if ret == true {
			t.Errorf("Add returned non-nil unexpectedly, ret=%v", ret)
		}
	}

	for i := 0; i < n; i++ {
		treap.Add(utils.V(i))
		ret := treap.FindEQ(utils.V(i))
		if ret == false {
			t.Errorf("Add returned nil unexpectedly")
		}
	}

	ret := treap.FindEQ(utils.V(n + 123))
	if ret == true {
		t.Errorf("Add returned nil unexpectedly")
	}
}

func TestFind(t *testing.T) {
	n := 10
	treap := New()

	for i := 0; i < n; i++ {
		ret := treap.Find(utils.V(i))
		if ret != nil {
			t.Errorf("non-nil unexpectedly, ret=%d", ret)
		}
	}

	for i := 0; i < n; i++ {
		treap.Add(utils.V(i))
		ret := treap.Find(utils.V(i))
		if ret == nil {
			t.Errorf("Find returned nil unexpectedly")
		}
	}

	ret := treap.Find(utils.V(n - 123))
	if *ret != utils.V(0) {
		t.Errorf("unexpectedly returned %v", *ret)
	}
}

func TestRemove(t *testing.T) {
	x := utils.V(-12345)
	treap := New()

	if treap.Remove(x) == true {
		t.Errorf("Add returned true unexpectedly")
	}

	treap.Add(x)
	if treap.Remove(x) == false {
		t.Errorf("Add returned false unexpectedly")
	}

	if treap.Find(x) != nil {
		t.Errorf("Add returned non-nil unexpectedly")
	}
}
