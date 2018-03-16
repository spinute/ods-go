package array_stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := [...]int{1, 2, 100, 0}

	for _, cap := range tests {
		as := New(cap)
		if as.n != 0 || as.cap != cap {
			t.Errorf("ArrayStack.New(%v): n = %v, cap = %v", cap, as.n, as.cap)
		}
	}
}

func TestPush(t *testing.T) {
	tests := [...]Value{1, 2, 1, -1, -2, -100, 0, 0}
	as := New(2)

	for i, v := range tests {
		as.Push(v)
		if as.n != i+1 {
			t.Errorf("as.n = %v at %v th push", as.n, i)
		}
	}
}

func TestPushMany(t *testing.T) {
	tests := make([]Value, 12345)
	as := New(0)

	for i, v := range tests {
		as.Push(v)
		if as.n != i+1 {
			t.Errorf("as.n = %v at %v th push", as.n, i)
		}
	}
}

func TestPushAndPop(t *testing.T) {
	tests := [...]Value{1, 2, 1, -1, -2, -100, 0, 0}
	as := New(1)

	for _, v := range tests {
		as.Push(v)
		if ret := as.Pop(); ret != v {
			t.Errorf("pushed %v, then %v was poped", v, ret)
		}
	}
}

func TestPushAndPop2(t *testing.T) {
	n := 123
	as := New(0)

	for i := 0; i < n; i++ {
		as.Push(Value(i + 1))
	}

	for i := 0; i < n; i++ {
		if ret, want := as.Pop(), Value(n-i); ret != want {
			t.Errorf("expect %d but returned %d", ret, want)
		}
	}
}
