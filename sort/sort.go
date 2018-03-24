package sort

import (
	"math/rand"

	"github.com/spinute/ods-go/utils"
)

func merge(a0, a1, a []utils.V) {
	i0 := 0
	i1 := 0
	for i := 0; i < len(a); i++ {
		if i0 == len(a0) {
			a[i] = a1[i1]
			i1++
		} else if i1 == len(a1) {
			a[i] = a0[i0]
			i0++
		} else if utils.Compare(a0[i0], a1[i1]) < 0 {
			a[i] = a0[i0]
			i0++
		} else {
			a[i] = a1[i1]
			i1++
		}
	}
}

func MergeSort(a []utils.V) []utils.V {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) / 2
	a0 := make([]utils.V, mid)
	copy(a0[0:mid], a[0:mid])
	a1 := make([]utils.V, len(a)-mid)
	copy(a1[0:len(a)-mid], a[mid:len(a)])
	MergeSort(a0)
	MergeSort(a1)
	merge(a0, a1, a)
	return a
}

func QuickSort(a []utils.V) []utils.V {
	quickSort(a, 0, len(a))
	return a
}

func quickSort(a []utils.V, i, n int) []utils.V {
	if n <= 1 {
		return a
	}
	x := a[i+rand.Int()%n]
	p := i - 1
	j := i
	q := i + n
	for j < q {
		comp := utils.Compare(a[j], x)
		if comp < 0 {
			p++
			a[j], a[p] = a[p], a[j]
			j++
		} else if comp > 0 {
			q--
			a[j], a[q] = a[q], a[j]
		} else {
			j++
		}
	}
	quickSort(a, i, p-i+1)
	quickSort(a, q, n-(q-i))
	return a
}
