package binary_heap

import (
	"github.com/spinute/ods-go/utils"
)

type BinaryHeap struct {
	n int
	a []utils.V
}

func New() *BinaryHeap {
	return &BinaryHeap{}
}

func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2*i + 2
}
func parent(i int) int {
	return (i - 1) / 2
}

func (bh *BinaryHeap) Add(x utils.V) bool {
	if bh.n+1 > len(bh.a) {
		bh.resize()
	}
	bh.a[bh.n] = x
	bh.n++
	bh.bubbleUp(bh.n - 1)
	return true
}
func (bh *BinaryHeap) bubbleUp(i int) {
	p := parent(i)
	for i > 0 && utils.Compare(bh.a[i], bh.a[p]) < 0 {
		tmp := bh.a[i]
		bh.a[i] = bh.a[p]
		bh.a[p] = tmp

		i = p
		p = parent(i)
	}
}

func (bh *BinaryHeap) resize() {
	cap_new := utils.Max(2*bh.n, 1)
	a_new := make([]utils.V, cap_new)
	copy(a_new, bh.a)
	bh.a = a_new
}

func (bh *BinaryHeap) Remove() utils.V {
	x := bh.a[0]
	bh.n--
	bh.a[0] = bh.a[bh.n]
	bh.trickleDown(0)
	if 3*bh.n < len(bh.a) {
		bh.resize()
	}
	return x
}
func (bh *BinaryHeap) trickleDown(i int) {
	for i >= 0 {
		j := -1
		r := right(i)
		if r < bh.n && utils.Compare(bh.a[r], bh.a[i]) < 0 {
			l := left(i)
			if utils.Compare(bh.a[l], bh.a[r]) < 0 {
				j = l
			} else {
				j = r
			}
		} else {
			l := left(i)
			if l < bh.n && utils.Compare(bh.a[l], bh.a[i]) < 0 {
				j = l
			}
		}
		if j >= 0 {
			tmp := bh.a[i]
			bh.a[i] = bh.a[j]
			bh.a[j] = tmp
		}
		i = j
	}
}
