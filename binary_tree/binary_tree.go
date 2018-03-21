package binary_tree

import "github.com/spinute/ods-go/utils"

type node struct {
	left, right, parent *node
}

func nodeNew() *node {
	return &node{}
}

type BinaryTree struct {
	r *node
}

func (bt BinaryTree) Depth(n *node) int {
	d := 0
	for n != bt.r {
		d++
		n = n.parent
	}
	return d
}

func size_subtree(n *node) int {
	if n == nil {
		return 0
	}
	return size_subtree(n.left) + size_subtree(n.right) + 1
}

func (bt BinaryTree) Size() int {
	return size_subtree(bt.r)
}

func height_subtree(n *node) int {
	if n == nil {
		return -1
	}
	return utils.Max(height_subtree(n.left), height_subtree(n.right)) + 1
}

func (bt BinaryTree) Height() int {
	return height_subtree(bt.r)
}
