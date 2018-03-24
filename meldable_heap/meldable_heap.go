package meldable_heap

import (
	"math/rand"

	"github.com/spinute/ods-go/utils"
)

type node struct {
	right, left, parent *node
	x                   utils.V
}

type MeldableHeap struct {
	r *node
	n int
}

func New() MeldableHeap {
	return MeldableHeap{}
}

func (mh *MeldableHeap) merge(h1, h2 *node) *node {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if utils.Compare(h1.x, h2.x) > 0 {
		return mh.merge(h2, h1)
	}
	if rand.Int()%2 == 0 {
		h1.left = mh.merge(h1.left, h2)
		if h1.left != nil {
			h1.left.parent = h1
		}
	} else {
		h1.right = mh.merge(h1.right, h2)
		if h1.right != nil {
			h1.right.parent = h1
		}
	}
	return h1
}

func (mh *MeldableHeap) Add(x utils.V) bool {
	u := &node{x: x}
	mh.r = mh.merge(u, mh.r)
	mh.r.parent = nil
	mh.n++
	return true
}

func (mh *MeldableHeap) Remove() utils.V {
	x := mh.r.x
	mh.r = mh.merge(mh.r.left, mh.r.right)
	if mh.r != nil {
		mh.r.parent = nil
	}
	mh.n--
	return x
}
