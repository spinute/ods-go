package skiplistsset

import (
	"math/rand"
	"strconv"

	"github.com/spinute/ods-go/utils"
)

type node struct {
	x      utils.V
	height int
	nexts  []*node
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
		height: h,
		x:      x,
		nexts:  make([]*node, h+1),
	}
}

type SkiplistSSet struct {
	sentinel  *node
	n, height int
	stack     [strconv.IntSize]*node
}

func New() *SkiplistSSet {
	ss := SkiplistSSet{
		sentinel: nodeNew(utils.V(0), strconv.IntSize),
	}
	ss.stack[0] = ss.sentinel
	return &ss
}

func (ss *SkiplistSSet) Find(x utils.V) *utils.V {
	p := ss.findPredNode(x)
	if p.nexts[0] == nil {
		return nil
	}
	return &p.nexts[0].x
}

func (ss *SkiplistSSet) findPredNode(x utils.V) *node {
	p := ss.sentinel
	r := ss.height
	for r >= 0 {
		for p.nexts[r] != nil && utils.Compare(p.nexts[r].x, x) < 0 {
			p = p.nexts[r]
		}
		r--
	}
	return p
}

func (ss *SkiplistSSet) Add(x utils.V) bool {
	p := ss.sentinel
	r := ss.height
	for r >= 0 {
		for p.nexts[r] != nil && utils.Compare(p.nexts[r].x, x) < 0 {
			p = p.nexts[r]
		}
		if p.nexts[r] != nil && utils.Compare(p.nexts[r].x, x) == 0 {
			return false
		}
		ss.stack[r] = p
		r--
	}
	h := pickHeight()
	newNode := nodeNew(x, h)
	for ss.height < newNode.height {
		ss.height++
		ss.stack[ss.height] = ss.sentinel
	}
	for i := 0; i <= newNode.height; i++ {
		newNode.nexts[i] = ss.stack[i].nexts[i]
		ss.stack[i].nexts[i] = newNode
	}
	ss.n++
	return true
}

func (ss *SkiplistSSet) Remove(x utils.V) bool {
	ret := false
	p := ss.sentinel
	r := ss.height
	for r >= 0 {
		for p.nexts[r] != nil && utils.Compare(p.nexts[r].x, x) < 0 {
			p = p.nexts[r]
		}
		if p.nexts[r] != nil && utils.Compare(p.nexts[r].x, x) == 0 {
			ret = true
			ss.n--
			p.nexts[r] = p.nexts[r].nexts[r]
			if p == ss.sentinel && p.nexts[r] == nil {
				ss.height--
			}
		}
		r--
	}
	return ret
}
