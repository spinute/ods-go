package scapegoat_tree

import (
	"testing"

	"github.com/spinute/ods-go/utils"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("ScapegoatTree.New().n = %v", ret)
	}
}

func TestAdd(t *testing.T) {
	n := 10
	st := New()

	for i := 0; i < n; i++ {
		ret := st.Add(utils.V(i))
		if st.n != i+1 {
			t.Errorf("st.n = %v at %v th Add", st.n, i+1)
		}
		if ret == false {
			t.Errorf("Add returned false unexpectedly")
		}
	}

	for i := 0; i < n; i++ {
		ret := st.Add(utils.V(i))
		if st.n != n {
			t.Errorf("st.n = %v", st.n)
		}
		if ret == true {
			t.Errorf("Add returned true unexpectedly")
		}
	}
}

func TestFindEQ(t *testing.T) {
	n := 10
	st := New()

	for i := 0; i < n; i++ {
		ret := st.FindEQ(utils.V(i))
		if ret == true {
			t.Errorf("Add returned non-nil unexpectedly, ret=%d", ret)
		}
	}

	for i := 0; i < n; i++ {
		st.Add(utils.V(i))
		ret := st.FindEQ(utils.V(i))
		if ret == false {
			t.Errorf("Add returned nil unexpectedly")
		}
	}

	ret := st.FindEQ(utils.V(n + 123))
	if ret == true {
		t.Errorf("Add returned nil unexpectedly")
	}
}

func TestFind(t *testing.T) {
	n := 10
	st := New()

	for i := 0; i < n; i++ {
		ret := st.Find(utils.V(i))
		if ret != nil {
			t.Errorf("non-nil unexpectedly, ret=%d", ret)
		}
	}

	for i := 0; i < n; i++ {
		st.Add(utils.V(i))
		ret := st.Find(utils.V(i))
		if ret == nil {
			t.Errorf("Find returned nil unexpectedly")
		}
	}

	ret := st.Find(utils.V(n - 123))
	if *ret != utils.V(0) {
		t.Errorf("unexpectedly returned %v", *ret)
	}
}

func TestRemove(t *testing.T) {
	x := utils.V(-12345)
	st := New()

	if st.Remove(x) == true {
		t.Errorf("Add returned true unexpectedly")
	}

	st.Add(x)
	if st.Remove(x) == false {
		t.Errorf("Add returned false unexpectedly")
	}

	if st.Find(x) != nil {
		t.Errorf("Add returned non-nil unexpectedly")
	}
}
