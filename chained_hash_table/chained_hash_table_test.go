package chained_hash_table

import (
	"testing"

	"github.com/spinute/ods-go/utils"
)

func TestNew(t *testing.T) {
	if ret := New().n; ret != 0 {
		t.Errorf("ChainedHashTable.New().n = %v", ret)
	}
}

func TestAdd(t *testing.T) {
	n := 10
	cht := New()

	for i := 0; i < n; i++ {
		ret := cht.Add(utils.V(i))
		if cht.n != i+1 {
			t.Errorf("cht.n = %v at %v th Add", cht.n, i+1)
		}
		if ret == false {
			t.Errorf("Add returned false unexpectedly")
		}
	}

	for i := 0; i < n; i++ {
		ret := cht.Add(utils.V(i))
		if cht.n != n {
			t.Errorf("cht.n = %v", cht.n)
		}
		if ret == true {
			t.Errorf("Add returned true unexpectedly")
		}
	}
}

func TestFind(t *testing.T) {
	n := 10
	cht := New()

	for i := 0; i < n; i++ {
		ret := cht.Find(utils.V(i))
		if ret != nil {
			t.Errorf("Add returned non-nil unexpectedly, ret=%d", ret)
		}
	}

	for i := 0; i < n; i++ {
		cht.Add(utils.V(i))
		ret := cht.Find(utils.V(i))
		if ret == nil {
			t.Errorf("Add returned nil unexpectedly")
		}
	}

	ret := cht.Find(utils.V(n + 123))
	if ret != nil {
		t.Errorf("Add returned nil unexpectedly")
	}
}

func TestRemove(t *testing.T) {
	x := utils.V(-12345)
	cht := New()

	if cht.Remove(x) == true {
		t.Errorf("Add returned true unexpectedly")
	}

	cht.Add(x)
	if cht.Remove(x) == false {
		t.Errorf("Add returned false unexpectedly")
	}

	if cht.Find(x) != nil {
		t.Errorf("Add returned non-nil unexpectedly")
	}
}
