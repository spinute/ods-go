package fast_array_stack

type Value int
type FastArrayStack struct {
	n, cap int
	buf    []Value
}

func New() FastArrayStack {
	return FastArrayStack{}
}

func (fas *FastArrayStack) Push(v Value) {
	fas.Add(fas.n, v)
}

func (fas *FastArrayStack) Pop() Value {
	return fas.Remove(fas.n - 1)
}

func (fas FastArrayStack) Get(i int) Value {
	return fas.buf[i]
}

func (fas *FastArrayStack) Set(i int, v Value) Value {
	ret := fas.buf[i]
	fas.buf[i] = v
	return ret
}

func (fas *FastArrayStack) Add(i int, v Value) {
	if fas.is_full() {
		fas.resize()
	}
	copy(fas.buf[i+1:fas.n+1], fas.buf[i:fas.n])
	fas.buf[i] = v
	fas.n++
}

func (fas *FastArrayStack) Remove(i int) Value {
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

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func (fas *FastArrayStack) resize() {
	fas.cap = max(2*fas.n, 1)
	buf_new := make([]Value, fas.cap)
	copy(buf_new, fas.buf)
	fas.buf = buf_new
}
