package fast_array_stack

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := [...]int{1, 2, 100, 0}

	for _, cap := range tests {
		fas := New(cap)
		if fas.n != 0 || fas.cap != cap {
			t.Errorf("FastArrayStack.New(%v): n = %v, cap = %v", cap, fas.n, fas.cap)
		}
	}
}

func TestPush(t *testing.T) {
	tests := [...]Value{1, 2, 1, -1, -2, -100, 0, 0}
	fas := New(2)

	for i, v := range tests {
		fas.Push(v)
		if fas.n != i+1 {
			t.Errorf("fas.n = %v at %v th push", fas.n, i)
		}
	}
}

func TestPushMany(t *testing.T) {
	tests := make([]Value, 12345)
	fas := New(0)

	for i, v := range tests {
		fas.Push(v)
		if fas.n != i+1 {
			t.Errorf("fas.n = %v at %v th push", fas.n, i)
		}
	}
}

func TestPushAndPop(t *testing.T) {
	tests := [...]Value{1, 2, 1, -1, -2, -100, 0, 0}
	fas := New(1)

	for _, v := range tests {
		fas.Push(v)
		if ret := fas.Pop(); ret != v {
			t.Errorf("pushed %v, then %v was poped", v, ret)
		}
	}
}

func TestPushAndPop2(t *testing.T) {
	n := 123
	fas := New(0)

	for i := 0; i < n; i++ {
		fas.Push(Value(i + 1))
	}

	for i := 0; i < n; i++ {
		if ret, want := fas.Pop(), Value(n-i); ret != want {
			t.Errorf("expect %d but returned %d", ret, want)
		}
	}
}
