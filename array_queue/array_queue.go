package array_queue

type Value int
type ArrayQueue struct {
	n, cap, i int
	buf       []Value
}

func New() ArrayQueue {
	return ArrayQueue{}
}

func (as *ArrayQueue) Add(v Value) {
	if as.is_full() {
		as.resize()
	}
	as.buf[(as.i+as.n)%as.cap] = v
	as.n++
}

func (as *ArrayQueue) Remove() Value {
	ret := as.buf[as.i]
	as.i = (as.i + 1) % as.cap
	as.n--
	if as.is_sparse() {
		as.resize()
	}
	return ret
}

func (as ArrayQueue) is_full() bool {
	return as.n == as.cap
}

func (as ArrayQueue) is_sparse() bool {
	return len(as.buf) >= 3*as.n
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func (as *ArrayQueue) resize() {
	cap_new := max(2*as.n, 1)
	buf_new := make([]Value, cap_new)
	for i := 0; i < as.n; i++ {
		buf_new[i] = as.buf[(i+as.i)%as.cap]
	}
	as.buf = buf_new
	as.cap = cap_new
	as.i = 0
}
