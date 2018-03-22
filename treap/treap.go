package treap

import (
	"math/rand"

	"github.com/spinute/ods-go/utils"
)

type node struct {
	left, right, parent *node
	x                   utils.V
	p                   int
}

type Treap struct {
	r *node
	n int
}

func New() *Treap {
	return &Treap{}
}

func (treap *Treap) FindEQ(x utils.V) bool {
	w := treap.r
	for w != nil {
		comp := utils.Compare(x, w.x)
		if comp < 0 {
			w = w.left
		} else if comp > 0 {
			w = w.right
		} else {
			return true
		}
	}
	return false
}

func (treap *Treap) Find(x utils.V) *utils.V {
	w := treap.r
	var ret *node = nil
	for w != nil {
		comp := utils.Compare(x, w.x)
		if comp < 0 {
			ret = w
			w = w.left
		} else if comp > 0 {
			w = w.right
		} else {
			return &w.x
		}
	}
	if ret == nil {
		return nil
	}
	return &ret.x
}

func (treap *Treap) Add(x utils.V) bool {
	var added bool
	newNode := &node{x: x, p: rand.Int()}
	lastNode := treap.findLast(x)

	if lastNode == nil {
		treap.r = newNode
		added = true
	} else {
		added = addChild(lastNode, newNode)
	}

	if added == true {
		treap.n++
		treap.bubble_up(newNode)
	}
	return added
}
func (treap *Treap) findLast(x utils.V) *node {
	w := treap.r
	last := w
	for w != nil {
		comp := utils.Compare(x, w.x)
		last = w
		if comp < 0 {
			w = w.left
		} else if comp > 0 {
			w = w.right
		} else {
			return w
		}
	}
	return last
}
func addChild(p, c *node) bool {
	comp := utils.Compare(c.x, p.x)
	if comp < 0 {
		p.left = c
	} else if comp > 0 {
		p.right = c
	} else {
		return false
	}
	c.parent = p
	return true
}
func (treap *Treap) bubble_up(u *node) {
	for u.parent != nil && u.parent.p > u.p {
		if u.parent.right == u {
			treap.rotateLeft(u.parent)
		} else {
			treap.rotateRight(u.parent)
		}
	}
	if u.parent == nil {
		treap.r = u
	}
}

func (treap *Treap) Remove(x utils.V) bool {
	u := treap.findLast(x)
	if u != nil && utils.Compare(u.x, x) == 0 {
		treap.trickleDown(u)
		treap.splice(u)
		return true
	}
	return false
}
func (treap *Treap) trickleDown(u *node) {
	for u.left != nil || u.right != nil {
		if u.left == nil {
			treap.rotateLeft(u)
		} else if u.right == nil {
			treap.rotateRight(u)
		} else if u.left.p < u.right.p {
			treap.rotateRight(u)
		} else {
			treap.rotateLeft(u)
		}
		if treap.r == u {
			treap.r = u.parent
		}
	}
}
func (treap *Treap) splice(u *node) {
	var c, p *node
	if u.left != nil {
		c = u.left
	} else {
		c = u.right
	}

	if treap.r == u {
		treap.r = c
		p = nil
	} else {
		p = u.parent
		if u == p.left {
			p.left = c
		} else {
			p.right = c
		}
	}

	if c != nil {
		c.parent = p
	}
}

func (treap *Treap) rotateLeft(u *node) {
	w := u.right
	w.parent = u.parent
	if w.parent != nil {
		if w.parent.left == u {
			w.parent.left = w
		} else {
			w.parent.right = w
		}
	}
	u.right = w.left
	if u.right != nil {
		u.right.parent = u
	}
	u.parent = w
	w.left = u
	if u == treap.r {
		treap.r = w
		treap.r.parent = nil
	}
}
func (treap *Treap) rotateRight(u *node) {
	w := u.left
	w.parent = u.parent
	if w.parent != nil {
		if w.parent.left == u {
			w.parent.left = w
		} else {
			w.parent.right = w
		}
	}
	u.left = w.right
	if u.left != nil {
		u.left.parent = u
	}
	u.parent = w
	w.right = u
	if u == treap.r {
		treap.r = w
		treap.r.parent = nil
	}
}
