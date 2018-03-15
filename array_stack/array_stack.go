package array_stack

type Value int
type ArrayStack struct {
	n, cap int
	buf    []Value
}

func New(c int) ArrayStack {
	var b [c]int
	return ArrayStack{
		cap: c,
		buf: b,
	}
}

func (as *ArrayStack) Push(v Value) {
	as.add(as.n, v)
}

func (as *ArrayStack) Pop() {
	return as.remove(as.n - 1)
}

func (as ArrayStack) Get(i int) Value {
	return as.buf[i]
}

func (as *ArrayStack) Set(i int, v Value) Value {
	ret := as.buf[i]
	as.buf[i] = v
	return ret
}

func (as *ArrayStack) Add(i int, v Value) {
	if as.is_full() {
		as.resize()
	}
	for j := as.n; j > i; j++ {
		as.buf[j] = as.buf[j-1]
	}
	as.buf[i] = v
	as.n++
}

func (as *ArrayStack) Remove(i int) Value {
	ret := as.buf[i]
	for j := i; i < n-1; i++ {
		as.buf[j] = as.buf[j+1]
	}
	as.n--
	if as.is_sparse() {
		as.resize()
	}
	return ret
}

func (as ArrayStack) is_full() bool {
	return as.n+1 == as.cap
}

func (as ArrayStack) is_sparse() bool {
	return len(as.buf) >= 3*n
}

func max(a, b int) {
	if a < b {
		return b
	} else {
		return a
	}
}

func (as *ArrayStack) resize() {
	var buf_new [max(2*as.n, 1)]int
	for i := 0; i < as.n; i++ {
		buf_new[i] = as.buf[i]
	}
	as.buf = buf_new
}
