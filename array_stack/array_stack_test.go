package array_stack

import (
	"github.com/spinute/ods-go/utils"
	"testing"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("ArrayStack.New().n = %v", ret)
	}
}

func TestPush(t *testing.T) {
	tests := []utils.V{1, 2, 1, -1, -2, -100, 0, 0}
	as := New()

	for i, v := range tests {
		as.Push(v)
		if as.n != i+1 {
			t.Errorf("as.n = %v at %v th push", as.n, i+1)
		}
	}
}

func TestPushMany(t *testing.T) {
	tests := make([]utils.V, 12345)
	as := New()

	for i, v := range tests {
		as.Push(v)
		if as.n != i+1 {
			t.Errorf("as.n = %v at %v th push", as.n, i+1)
		}
	}
}

func TestPushAndPop(t *testing.T) {
	tests := []utils.V{1, 2, 1, -1, -2, -100, 0, 0}
	as := New()

	for _, v := range tests {
		as.Push(v)
		if ret := as.Pop(); ret != v {
			t.Errorf("pushed %v, then %v was poped", v, ret)
		}
	}
}

func TestPushAndPop2(t *testing.T) {
	n := 123
	as := New()

	for i := 0; i < n; i++ {
		as.Push(utils.V(i + 1))
	}

	for i := 0; i < n; i++ {
		if ret, want := as.Pop(), utils.V(n-i); ret != want {
			t.Errorf("expect %d but returned %d", ret, want)
		}
	}
}
