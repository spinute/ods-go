package rootish_array_stack

import (
	"github.com/spinute/ods-go/utils"
	"math"
)

type RootishArrayStack struct {
	n      int
	blocks [][]utils.V
}

func New() RootishArrayStack {
	return RootishArrayStack{}
}

func i2b(i int) int {
	db := (-3.0 + math.Sqrt(float64(9+8*i))) / 2.0
	return int(math.Ceil(db))
}

func (ras RootishArrayStack) Get(i int) utils.V {
	b := i2b(i)
	j := i - b*(b+1)/2
	return ras.blocks[b][j]
}

func (ras *RootishArrayStack) Set(i int, v utils.V) utils.V {
	b := i2b(i)
	j := i - b*(b+1)/2
	ret := ras.blocks[b][j]
	ras.blocks[b][j] = v
	return ret
}

func (ras RootishArrayStack) is_full() bool {
	n_b := len(ras.blocks)
	return n_b*(n_b+1)/2 == ras.n
}
func (ras *RootishArrayStack) grow() {
	ras.blocks = append(ras.blocks, make([]utils.V, len(ras.blocks)+1))
}

func (ras *RootishArrayStack) Add(i int, v utils.V) {
	if ras.is_full() {
		ras.grow()
	}
	for j := ras.n; j > i; j-- {
		ras.Set(j, ras.Get(j-1))
	}
	ras.Set(i, v)
	ras.n++
}

func (ras RootishArrayStack) is_sparse() bool {
	n_b := len(ras.blocks)
	return n_b > 0 && (n_b-2)*(n_b-1)/2 >= ras.n
}
func (ras *RootishArrayStack) shrink() {
	// go slice will not shrink ??
}

func (ras *RootishArrayStack) Remove(i int) utils.V {
	ret := ras.Get(i)
	for j := i; j < ras.n-1; j++ {
		ras.Set(j, ras.Get(j+1))
	}
	ras.n--
	n_b := len(ras.blocks)
	if (n_b-2)*(n_b-1)/2 >= ras.n {
		ras.shrink()
	}
	return ret
}
