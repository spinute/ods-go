package dllist

import (
	"github.com/spinute/ods-go/utils"
	"testing"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("DLList.New().n = %v", ret)
	}
}

func TestAdd(t *testing.T) {
	n := 10
	dll := New()

	for i := 0; i < n; i++ {
		dll.Add(utils.V(i))
		if dll.n != i+1 {
			t.Errorf("dll.n = %v at %v th Add", dll.n, i+1)
		}
	}
}

func TestPush(t *testing.T) {
	n := 10
	dll := New()

	for i := 0; i < n; i++ {
		dll.Push(utils.V(i))
		if dll.n != i+1 {
			t.Errorf("dll.n = %v at %v th Push", dll.n, i+1)
		}
	}
}

func TestPushAndPop(t *testing.T) {
	n := 10
	dll := New()
	for i := 0; i < n; i++ {
		dll.Push(utils.V(i))
	}

	for i := 0; i < n; i++ {
		if ret := dll.Pop(); ret != utils.V(n-1-i) {
			t.Errorf("dll.Pop() = %v", i, ret)
		}
	}
}

func TestAddAndRemove(t *testing.T) {
	n := 10
	dll := New()
	for i := 0; i < n; i++ {
		dll.Add(utils.V(i))
	}

	for i := 0; i < n; i++ {
		if ret := dll.Remove(); ret != utils.V(i) {
			t.Errorf("dll.Remove() = %v", i, ret)
		}
	}
}
