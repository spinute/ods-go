package binary_search_tree

import "github.com/spinute/ods-go/utils"

type node struct {
	left, right, parent *node
	x                   utils.V
}

type BinarySearchTree struct {
	r *node
	n int
}

func New() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (bst *BinarySearchTree) FindEQ(x utils.V) bool {
	w := bst.r
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

func (bst *BinarySearchTree) Find(x utils.V) *utils.V {
	w := bst.r
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

func (bst *BinarySearchTree) Add(x utils.V) bool {
	var ret bool
	lastNode := bst.findLast(x)
	newNode := node{x: x}
	if lastNode == nil {
		bst.r = &newNode
		ret = true
	} else {
		ret = addChild(lastNode, &newNode)
	}
	if ret == true {
		bst.n++
	}
	return ret
}
func (bst *BinarySearchTree) findLast(x utils.V) *node {
	w := bst.r
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

func (bst *BinarySearchTree) Remove(x utils.V) bool {
	node := bst.findLast(x)
	if node != nil && utils.Compare(x, node.x) == 0 {
		bst.removeNode(node)
		return true
	}
	return false
}
func (bst *BinarySearchTree) removeNode(node *node) {
	if node.left == nil || node.right == nil {
		bst.splice(node)
	} else {
		minInRight := node.right
		for minInRight.left != nil {
			minInRight = minInRight.left
		}
		node.x = minInRight.x
		bst.splice(minInRight)
	}
	bst.n--
}
func (bst *BinarySearchTree) splice(n *node) {
	var c, p *node
	if n.left != nil {
		c = n.left
	} else {
		c = n.right
	}

	if bst.r == n {
		bst.r = c
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
