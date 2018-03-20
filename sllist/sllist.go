package sllist

import (
	"github.com/spinute/ods-go/utils"
)

type node struct {
	x    utils.V
	next *node
}

func nodeNew(x utils.V) *node {
	return &node{
		x: x,
	}
}

type SLList struct {
	n          int
	head, tail *node
}

func New() *SLList {
	return new(SLList)
}

func (l *SLList) Size() int {
	return l.n
}

func (l *SLList) Push(x utils.V) {
	new_node := nodeNew(x)
	new_node.next = l.head
	l.head = new_node
	if l.n == 0 {
		l.tail = new_node
	}
	l.n++
}

func (l *SLList) Pop() utils.V {
	if l.n == 0 {
		return 0 // FIXME: handle error
	}

	ret := l.head.x
	l.head = l.head.next

	if l.n == 1 {
		l.tail = nil
	}
	l.n--

	return ret
}

func (l *SLList) Remove() utils.V {
	return l.Pop()
}

func (l *SLList) Add(x utils.V) {
	new_node := nodeNew(x)
	if l.n == 0 {
		l.head = new_node
	} else {
		l.tail.next = new_node
	}

	l.tail = new_node
	l.n++
}
