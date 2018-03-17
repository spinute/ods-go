package array_stack

import (
	"github.com/spinute/ods-go/utils"
)

type ArrayStack struct {
	n, cap int
	buf    []utils.V
}

func New() ArrayStack {
	return ArrayStack{}
}

func (as ArrayStack) Size() int {
	return as.n
}

func (as *ArrayStack) Push(v utils.V) {
	as.Add(as.n, v)
}

func (as *ArrayStack) Pop() utils.V {
	return as.Remove(as.n - 1)
}

func (as ArrayStack) Get(i int) utils.V {
	return as.buf[i]
}

func (as *ArrayStack) Set(i int, v utils.V) utils.V {
	ret := as.buf[i]
	as.buf[i] = v
	return ret
}

func (as *ArrayStack) Add(i int, v utils.V) {
	if as.is_full() {
		as.resize()
	}
	for j := as.n; j > i; j-- {
		as.buf[j] = as.buf[j-1]
	}
	as.buf[i] = v
	as.n++
}

func (as *ArrayStack) Remove(i int) utils.V {
	ret := as.buf[i]
	for j := i; j < as.n-1; j++ {
		as.buf[j] = as.buf[j+1]
	}
	as.n--
	if as.is_sparse() {
		as.resize()
	}
	return ret
}

func (as ArrayStack) is_full() bool {
	return as.n == as.cap
}

func (as ArrayStack) is_sparse() bool {
	return len(as.buf) >= 3*as.n
}

func (as *ArrayStack) resize() {
	as.cap = utils.Max(2*as.n, 1)
	buf_new := make([]utils.V, as.cap)
	for i := 0; i < as.n; i++ {
		buf_new[i] = as.buf[i]
	}
	as.buf = buf_new
}
