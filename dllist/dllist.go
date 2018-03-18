package dllist

import (
	"github.com/spinute/ods-go/utils"
)

type node struct {
	x          utils.V
	next, prev *node
}

func nodeNew(x utils.V) *node {
	return &node{
		x: x,
	}
}

type DLList struct {
	n     int
	dummy *node
}

func New() *DLList {
	dummy := new(node)
	dummy.next = dummy
	dummy.prev = dummy
	return &DLList{
		dummy: dummy,
	}
}

func (l *DLList) get_node(i int) *node {
	var p *node
	if i < l.n/2 {
		p = l.dummy.next
		for j := 0; j < i; j++ {
			p = p.next
		}
	} else {
		p = l.dummy
		for j := l.n; j > i; j-- {
			p = p.prev
		}
	}
	return p
}

func (l *DLList) Get(i int) utils.V {
	return l.get_node(i).x
}
func (l *DLList) Set(i int, v utils.V) utils.V {
	p := l.get_node(i)
	ret := p.x
	p.x = v
	return ret
}

func (l *DLList) add_before(p *node, x utils.V) {
	new_node := nodeNew(x)
	new_node.prev = p.prev
	new_node.next = p
	new_node.prev.next = new_node
	new_node.next.prev = new_node
	l.n++
}

func (l *DLList) add_i(i int, x utils.V) {
	l.add_before(l.get_node(i), x)
}

func (l *DLList) remove_i(i int) utils.V {
	p := l.get_node(i)
	ret := p.x
	p.prev.next = p.next
	p.next.prev = p.prev
	l.n--
	return ret
}

func (l *DLList) Push(x utils.V) {
	l.add_i(l.n, x)
}

func (l *DLList) Add(x utils.V) {
	l.add_i(l.n, x)
}

func (l *DLList) Pop() utils.V {
	return l.remove_i(l.n - 1)
}

func (l *DLList) Remove() utils.V {
	return l.remove_i(0)
}
