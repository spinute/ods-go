package fast_array_stack

import "github.com/spinute/ods-go/utils"

type FastArrayStack struct {
	n, cap int
	buf    []utils.V
}

func New() FastArrayStack {
	return FastArrayStack{}
}

func (fas *FastArrayStack) Push(v utils.V) {
	fas.Add(fas.n, v)
}

func (fas *FastArrayStack) Pop() utils.V {
	return fas.Remove(fas.n - 1)
}

func (fas FastArrayStack) Get(i int) utils.V {
	return fas.buf[i]
}

func (fas *FastArrayStack) Set(i int, v utils.V) utils.V {
	ret := fas.buf[i]
	fas.buf[i] = v
	return ret
}

func (fas *FastArrayStack) Add(i int, v utils.V) {
	if fas.is_full() {
		fas.resize()
	}
	copy(fas.buf[i+1:fas.n+1], fas.buf[i:fas.n])
	fas.buf[i] = v
	fas.n++
}

func (fas *FastArrayStack) Remove(i int) utils.V {
	ret := fas.buf[i]
	copy(fas.buf[i:fas.n-1], fas.buf[i+1:fas.n])
	fas.n--
	if fas.is_sparse() {
		fas.resize()
	}
	return ret
}

func (fas FastArrayStack) is_full() bool {
	return fas.n == fas.cap
}

func (fas FastArrayStack) is_sparse() bool {
	return len(fas.buf) >= 3*fas.n
}

func (fas *FastArrayStack) resize() {
	fas.cap = utils.Max(2*fas.n, 1)
	buf_new := make([]utils.V, fas.cap)
	copy(buf_new, fas.buf)
	fas.buf = buf_new
}
