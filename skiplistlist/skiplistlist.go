package skiplistlist

import (
	"math/rand"
	"strconv"

	"github.com/spinute/ods-go/utils"
)

type node struct {
	x       utils.V
	height  int
	lengths []int
	nexts   []*node
}

func pickHeight() int {
	r := rand.Int()
	h := 0
	for r&1 != 0 {
		h++
		r >>= 1
	}
	return h
}

func nodeNew(x utils.V, h int) *node {
	return &node{
		height:  h,
		x:       x,
		lengths: make([]int, h+1),
		nexts:   make([]*node, h+1),
	}
}

type SkiplistList struct {
	sentinel  *node
	n, height int
	stack     [strconv.IntSize]*node
}

func New() *SkiplistList {
	ss := SkiplistList{
		sentinel: nodeNew(utils.V(0), strconv.IntSize),
	}
	ss.stack[0] = ss.sentinel
	return &ss
}

func (ss *SkiplistList) findPred(i int) *node {
	u := ss.sentinel
	r := ss.height
	j := -1
	for r >= 0 {
		for u.nexts[r] != nil && j+u.lengths[r] < i {
			j += u.lengths[r]
			u = u.nexts[r]
		}
		r--
	}
	return u
}

func (ss *SkiplistList) Get(i int) utils.V {
	return ss.findPred(i).nexts[0].x
}

func (ss *SkiplistList) Set(i int, x utils.V) utils.V {
	u := ss.findPred(i).nexts[0]
	y := u.x
	u.x = x
	return y
}

func (ss *SkiplistList) Add(i int, x utils.V) {
	w := nodeNew(x, pickHeight())
	if w.height > ss.height {
		ss.height = w.height
	}
	ss.addNode(i, w)
}

func (ss *SkiplistList) addNode(i int, w *node) *node {
	u := ss.sentinel
	k := w.height
	r := ss.height
	j := -1
	for r >= 0 {
		for u.nexts[r] != nil && j+u.lengths[r] < i {
			j += u.lengths[r]
			u = u.nexts[r]
		}
		u.lengths[r]++
		if r <= k {
			w.nexts[r] = u.nexts[r]
			u.nexts[r] = w
			w.lengths[r] = u.lengths[r] - (i - j)
			u.lengths[r] = i - j
		}
		r--
	}
	ss.n++
	return u
}

func (ss *SkiplistList) Remove(i int) utils.V {
	// FIXME: return error when i th value does not exist
	ret := utils.V(-12345)
	orig := ret

	u := ss.sentinel
	r := ss.height
	j := -1
	for r >= 0 {
		for u.nexts[r] != nil && j+u.lengths[r] < i {
			j += u.lengths[r]
			u = u.nexts[r]
		}
		u.lengths[r]--
		if j+u.lengths[r]+1 == i && u.nexts[r] != nil {
			ret = u.nexts[r].x
			u.lengths[r] += u.nexts[r].lengths[r]
			u.nexts[r] = u.nexts[r].nexts[r]
			if u == ss.sentinel && u.nexts[r] == nil {
				ss.height--
			}
		}
		r--
	}

	// FIXME: return error when i th value does not exist
	if ret != orig {
		ss.n--
	}

	return ret
}
