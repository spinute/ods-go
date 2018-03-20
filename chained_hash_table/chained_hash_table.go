package chained_hash_table

import (
	"github.com/spinute/ods-go/array_stack"
	"github.com/spinute/ods-go/utils"
	"math/rand"
)

type ChainedHashTable struct {
	n    int
	t    []array_stack.ArrayStack
	z, d uint64
}

func New() *ChainedHashTable {
	return &ChainedHashTable{
		d: 1,
		t: make([]array_stack.ArrayStack, 1<<1),
		z: rand.Uint64() | uint64(1),
	}
}

func (cht *ChainedHashTable) hash(x utils.V) uint64 {
	return (cht.z * uint64(x)) >> (64 - cht.d)
}

func (cht *ChainedHashTable) Find(x utils.V) *utils.V {
	h := cht.hash(x)
	for i := 0; i < cht.t[h].Size(); i++ {
		if x == cht.t[h].Get(i) {
			ret := cht.t[h].Get(i)
			return &ret
		}
	}
	return nil
}
func (cht *ChainedHashTable) Add(x utils.V) bool {
	if cht.Find(x) != nil {
		return false
	}
	if cht.n+1 > len(cht.t) {
		cht.resize()
	}
	cht.t[cht.hash(x)].Push(x)
	cht.n++
	return true
}

func (cht *ChainedHashTable) Remove(x utils.V) bool {
	j := cht.hash(x)
	for i := 0; i < cht.t[j].Size(); i++ {
		ret := cht.t[j].Get(i)
		if x == ret {
			cht.t[j].Remove(i)
			cht.n--
			return true
		}
	}
	return false
}

func (cht *ChainedHashTable) resize() {
	dNew := uint64(1)
	for 1<<dNew <= cht.n {
		dNew++
	}
	cht.d = dNew
	tNew := make([]array_stack.ArrayStack, 1<<cht.d)
	for i := 0; i < len(cht.t); i++ {
		t := cht.t[i]
		for j := 0; j < t.Size(); j++ {
			x := t.Get(j)
			tNew[cht.hash(x)].Push(x)
		}
	}
	cht.t = tNew
}
