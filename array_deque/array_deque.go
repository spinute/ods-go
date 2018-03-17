package array_deque

import "github.com/spinute/ods-go/utils"

type ArrayDeque struct {
	n, cap, i int
	buf       []utils.V
}

func New() ArrayDeque {
	return ArrayDeque{}
}

func (ad ArrayDeque) buf_i(i int) int {
	return (ad.i + i) % ad.cap
}
func (ad ArrayDeque) buf_get(i int) utils.V {
	return ad.buf[ad.buf_i(i)]
}
func (ad *ArrayDeque) buf_set(i int, v utils.V) {
	ad.buf[ad.buf_i(i)] = v
}

func (ad ArrayDeque) Get(i int) utils.V {
	return ad.buf_get(i)
}

func (ad *ArrayDeque) Set(i int, v utils.V) utils.V {
	ret := ad.buf_get(i)
	ad.buf_set(i, v)
	return ret
}

func mod_pred(i, mod int) int {
	if i == 0 {
		return mod - 1
	} else {
		return i - 1
	}
}

func (ad *ArrayDeque) Add(i int, v utils.V) {
	if ad.is_full() {
		ad.resize()
	}

	if i < ad.n/2 {
		ad.i = mod_pred(ad.i, ad.cap)
		for j := 0; j < i; j++ {
			ad.buf_set(j, ad.buf_get(j+1))
		}
	} else {
		for j := ad.n; j > i; j-- {
			ad.buf_set(j, ad.buf_get(j-1))
		}
	}
	ad.buf_set(i, v)
	ad.n++
}

func mod_succ(i, mod int) int {
	return (i + 1) % mod
}

func (ad *ArrayDeque) Remove(i int) utils.V {
	ret := ad.buf_get(i)
	if i < ad.n/2 {
		for j := i; j > 0; j-- {
			ad.buf_set(j, ad.buf_get(j-1))
		}
		ad.i = mod_succ(ad.i, ad.cap)
	} else {
		for j := i; j < ad.n-1; j++ {
			ad.buf_set(j, ad.buf_get(j+1))
		}
	}
	ad.n--
	if ad.is_sparse() {
		ad.resize()
	}
	return ret
}

func (ad ArrayDeque) is_full() bool {
	return ad.n == ad.cap
}

func (ad ArrayDeque) is_sparse() bool {
	return len(ad.buf) >= 3*ad.n
}

func (ad *ArrayDeque) resize() {
	cap_new := utils.Max(2*ad.n, 1)
	buf_new := make([]utils.V, cap_new)
	for i := 0; i < ad.n; i++ {
		buf_new[i] = ad.buf_get(i)
	}
	ad.i = 0
	ad.buf = buf_new
	ad.cap = cap_new
}
