package linear_hash_table

import (
	"math/rand"

	"github.com/spinute/ods-go/utils"
)

type LinearHashTable struct {
	n, q int
	t    []*utils.V
	d, z uint64
	del  *utils.V
}

func New() *LinearHashTable {
	return &LinearHashTable{
		d:   1,
		t:   make([]*utils.V, 1<<1),
		z:   rand.Uint64() | uint64(1),
		del: new(utils.V),
	}
}

func (lht *LinearHashTable) hash(x utils.V) uint64 {
	return (lht.z * uint64(x)) >> (64 - lht.d)
}

func (lht *LinearHashTable) next_ind(i uint64) uint64 {
	return (i + 1) % uint64(len(lht.t))
}

func (lht *LinearHashTable) Find(x utils.V) *utils.V {
	h := lht.hash(x)
	for i := h; lht.t[i] != nil; i = lht.next_ind(i) {
		if lht.t[i] != lht.del && *lht.t[i] == x {
			return lht.t[i]
		}
	}
	return nil
}
func (lht *LinearHashTable) Add(x utils.V) bool {
	if lht.Find(x) != nil {
		return false
	}

	if 2*(lht.q+1) > len(lht.t) {
		lht.resize()
	}
	i := lht.hash(x)
	for lht.t[i] != nil && lht.t[i] != lht.del {
		i = lht.next_ind(i)
	}
	if lht.t[i] == nil {
		lht.q++
	}
	lht.n++
	lht.t[i] = &x
	return true
}

func (lht *LinearHashTable) Remove(x utils.V) bool {
	i := lht.hash(x)
	for lht.t[i] != nil {
		y := lht.t[i]
		if y != lht.del && x == *y {
			lht.t[i] = lht.del
			lht.n--
			if 8*lht.n < len(lht.t) {
				lht.resize()
			}
			return true
		}
		i = lht.next_ind(i)
	}
	return false
}

func (lht *LinearHashTable) resize() {
	dNew := uint64(1)
	for 1<<dNew < 3*lht.n {
		dNew++
	}
	lht.d = dNew
	lht.q = lht.n
	tNew := make([]*utils.V, 1<<lht.d)
	for i := 0; i < len(lht.t); i++ {
		if x := lht.t[i]; x != nil && x != lht.del {
			j := lht.hash(*x)
			for tNew[j] != nil {
				j = lht.next_ind(j)
			}
			tNew[j] = x
		}
	}
	lht.t = tNew
}
