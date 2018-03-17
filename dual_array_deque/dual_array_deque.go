package dual_array_deque

import (
	"github.com/spinute/ods-go/array_stack"
	"github.com/spinute/ods-go/utils"
)

type DualArrayDeque struct {
	n           int
	front, back array_stack.ArrayStack
}

func New() DualArrayDeque {
	return DualArrayDeque{
		front: array_stack.New(),
		back:  array_stack.New(),
	}
}

func (dad DualArrayDeque) Get(i int) utils.V {
	if front_size := dad.front.Size(); i < front_size {
		return dad.front.Get(front_size - i - 1)
	} else {
		return dad.back.Get(i - front_size)
	}
}

func (dad *DualArrayDeque) Set(i int, v utils.V) utils.V {
	if front_size := dad.front.Size(); i < front_size {
		return dad.front.Set(front_size-i-1, v)
	} else {
		return dad.back.Set(i-front_size, v)
	}
}

func (dad *DualArrayDeque) Add(i int, v utils.V) {
	if front_size := dad.front.Size(); i < front_size {
		dad.front.Add(front_size-i, v)
	} else {
		dad.back.Add(i-front_size, v)
	}
	dad.balance()
	dad.n++
}

func (dad *DualArrayDeque) Remove(i int) utils.V {
	var ret utils.V
	if front_size := dad.front.Size(); i < front_size {
		ret = dad.front.Remove(front_size - i - 1)
	} else {
		ret = dad.back.Remove(i - front_size)
	}
	dad.balance()
	dad.n--
	return ret
}

func (dad *DualArrayDeque) balance() {
	front_size := dad.front.Size()
	back_size := dad.back.Size()

	if 3*front_size < back_size || 3*back_size < front_size {
		n := front_size + back_size

		front_size_new := n / 2
		front_new := array_stack.New()
		for i := front_size_new - 1; i >= 0; i-- {
			front_new.Push(dad.Get(i))
		}

		back_size_new := n - front_size_new
		back_new := array_stack.New()
		for i := 0; i < back_size_new; i++ {
			back_new.Push(dad.Get(front_size_new + i))
		}

		dad.front = front_new
		dad.back = back_new
	}
}
