package skiplistsset

import (
	"testing"

	"github.com/spinute/ods-go/utils"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("SkiplistSSet.New().n = %v", ret)
	}
}

func TestAdd(t *testing.T) {
	n := 10
	ss := New()

	for i := 0; i < n; i++ {
		ret := ss.Add(utils.V(i))
		if ss.n != i+1 {
			t.Errorf("ss.n = %v at %v th Add", ss.n, i+1)
		}
		if ret == false {
			t.Errorf("Add returned false unexpectedly")
		}
	}

	for i := 0; i < n; i++ {
		ret := ss.Add(utils.V(i))
		if ss.n != n {
			t.Errorf("ss.n = %v", ss.n)
		}
		if ret == true {
			t.Errorf("Add returned true unexpectedly")
		}
	}
}

func TestFind(t *testing.T) {
	n := 10
	ss := New()

	for i := 0; i < n; i++ {
		ret := ss.Find(utils.V(i))
		if ret != nil {
			t.Errorf("Add returned non-nil unexpectedly, ret=%d", ret)
		}
	}

	for i := 0; i < n; i++ {
		ss.Add(utils.V(i))
		ret := ss.Find(utils.V(i))
		if ret == nil {
			t.Errorf("Add returned nil unexpectedly")
		}
	}

	ret := ss.Find(utils.V(n + 123))
	if ret != nil {
		t.Errorf("Add returned nil unexpectedly")
	}
}

func TestRemove(t *testing.T) {
	x := utils.V(-12345)
	ss := New()

	if ss.Remove(x) == true {
		t.Errorf("Add returned true unexpectedly")
	}

	ss.Add(x)
	if ss.Remove(x) == false {
		t.Errorf("Add returned false unexpectedly")
	}

	if ss.Find(x) != nil {
		t.Errorf("Add returned non-nil unexpectedly")
	}
}
