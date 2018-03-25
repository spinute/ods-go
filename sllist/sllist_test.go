package sllist

import (
	"github.com/spinute/ods-go/utils"
	"testing"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("SLList.New().n = %v", ret)
	}
}

func TestAdd(t *testing.T) {
	n := 10
	sll := New()

	for i := 0; i < n; i++ {
		sll.Add(utils.V(i))
		if sll.n != i+1 {
			t.Errorf("sll.n = %v at %v th Add", sll.n, i+1)
		}
	}
}

func TestPush(t *testing.T) {
	n := 10
	sll := New()

	for i := 0; i < n; i++ {
		sll.Push(utils.V(i))
		if sll.n != i+1 {
			t.Errorf("sll.n = %v at %v th Push", sll.n, i+1)
		}
	}
}

func TestPushAndPop(t *testing.T) {
	n := 10
	sll := New()
	for i := 0; i < n; i++ {
		sll.Push(utils.V(i))
	}

	for i := 0; i < n; i++ {
		if ret := sll.Pop(); ret != utils.V(n-1-i) {
			t.Errorf("%v th sll.Pop() = %v", i, ret)
		}
	}
}

func TestAddAndRemove(t *testing.T) {
	n := 10
	sll := New()
	for i := 0; i < n; i++ {
		sll.Add(utils.V(i))
	}

	for i := 0; i < n; i++ {
		if ret := sll.Remove(); ret != utils.V(i) {
			t.Errorf("%v th sll.Remove() = %v", i, ret)
		}
	}
}
