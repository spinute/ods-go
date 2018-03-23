package scapegoat_tree

import (
	"math"

	"github.com/spinute/ods-go/utils"
)

type node struct {
	left, right, parent *node
	x                   utils.V
}

type ScapegoatTree struct {
	r    *node
	n, q int
}

func New() *ScapegoatTree {
	return &ScapegoatTree{}
}

func (st *ScapegoatTree) FindEQ(x utils.V) bool {
	w := st.r
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

func (st *ScapegoatTree) Find(x utils.V) *utils.V {
	w := st.r
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

func size_subtree(n *node) int {
	if n == nil {
		return 0
	}
	return size_subtree(n.left) + size_subtree(n.right) + 1
}
func (st *ScapegoatTree) rebuild(u *node) {
	sz := size_subtree(u)
	p := u.parent
	a := make([]*node, sz)
	packIntoArray(u, a, 0)
	if p == nil {
		st.r = st.buildBalanced(a, 0, sz)
		st.r.parent = nil
	} else if p.right == u {
		p.right = st.buildBalanced(a, 0, sz)
		p.right.parent = p
	} else {
		p.left = st.buildBalanced(a, 0, sz)
		p.left.parent = p
	}
}
func packIntoArray(u *node, a []*node, i int) int {
	if u == nil {
		return i
	}
	i = packIntoArray(u.left, a, i)
	a[i] = u
	i++
	return packIntoArray(u.right, a, i)
}
func (st *ScapegoatTree) buildBalanced(a []*node, i, ns int) *node {
	if ns == 0 {
		return nil
	}
	m := ns / 2
	a[i+m].left = st.buildBalanced(a, i, m)
	if a[i+m].left != nil {
		a[i+m].left.parent = a[i+m]
	}
	a[i+m].right = st.buildBalanced(a, i+m+1, ns-m-1)
	if a[i+m].right != nil {
		a[i+m].right.parent = a[i+m]
	}
	return a[i+m]
}

func log32(n int) int {
	return int(math.Log(float64(n)) / math.Log(3.0/2))
}
func (st *ScapegoatTree) Add(x utils.V) bool {
	u := &node{x: x}
	d := st.addWithDepth(u)
	if d > log32(st.q) {
		w := u.parent
		a := size_subtree(w)
		b := size_subtree(w.parent)
		for 3*a <= 2*b {
			w = w.parent
			a = size_subtree(w)
			b = size_subtree(w.parent)
		}
		st.rebuild(w.parent)
	} else if d < 0 {
		return false
	}
	return true
}
func (st *ScapegoatTree) addWithDepth(u *node) int {
	w := st.r
	if w == nil {
		st.r = u
		st.n++
		st.q++
		return 0
	}
	done := false
	d := 0
	for !done {
		res := utils.Compare(u.x, w.x)
		if res < 0 {
			if w.left == nil {
				w.left = u
				u.parent = w
				done = true
			} else {
				w = w.left
			}
		} else if res > 0 {
			if w.right == nil {
				w.right = u
				u.parent = w
				done = true
			}
			w = w.right
		} else {
			return -1
		}
		d++
	}
	st.n++
	st.q++
	return d
}

func (st *ScapegoatTree) findLast(x utils.V) *node {
	w := st.r
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

func (st *ScapegoatTree) Remove(x utils.V) bool {
	if st.remove_value(x) {
		if 2*st.n < st.q {
			if st.r != nil {
				st.rebuild(st.r)
			}
			st.q = st.n
		}
		return true
	}
	return false
}
func (st *ScapegoatTree) remove_value(x utils.V) bool {
	node := st.findLast(x)
	if node != nil && utils.Compare(x, node.x) == 0 {
		st.removeNode(node)
		return true
	}
	return false
}
func (st *ScapegoatTree) removeNode(node *node) {
	if node.left == nil || node.right == nil {
		st.splice(node)
	} else {
		minInRight := node.right
		for minInRight.left != nil {
			minInRight = minInRight.left
		}
		node.x = minInRight.x
		st.splice(minInRight)
	}
	st.n--
}
func (st *ScapegoatTree) splice(n *node) {
	var c, p *node
	if n.left != nil {
		c = n.left
	} else {
		c = n.right
	}

	if st.r == n {
		st.r = c
		p = nil
	} else {
		p = n.parent
		if n == p.left {
			p.left = c
		} else {
			p.right = c
		}
	}

	if c != nil {
		c.parent = p
	}
}
