package sort

import (
	"cmp"
	"math/rand"

	"github.com/Savvelius/graph-algs/heap"
)

// swaps elements at indexes i1 and i2
func Swap[T any](slice []T, i1, i2 int) {
	slice[i1], slice[i2] = slice[i2], slice[i1]
}

// sort in ascending order
func Choice[T cmp.Ordered](slice []T) {
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[i] {
				Swap(slice, i, j)
			}
		}
	}
}

func Bubble[T cmp.Ordered](slice []T) {
	for k := 0; k < len(slice); k++ {
		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				Swap(slice, i, i+1)
			}
		}
	}
}

func Insertion[T cmp.Ordered](slice []T) {
	for k := 1; k < len(slice); k++ {
		for i := k; i > 0 && slice[i] < slice[i-1]; i-- {
			Swap(slice, i, i-1)
		}
	}
}

func Heap[T cmp.Ordered](slice []T) []T {
	hp := heap.New(func(t1, t2 T) bool {
		return t1 < t2
	}, len(slice))

	for _, elem := range slice {
		hp.Insert(elem)
	}

	return hp.ToSlice()
}

// asserts that elements in slice are in range [l, r]
func Count(slice []int, l, r int) {
	counts := make([]int, r-l+1)

	for _, elem := range slice {
		counts[elem-l]++
	}

	counted := 0
	for i, n := range counts {
		for j := counted; j < counted+n; j++ {
			slice[j] = i + l
		}
		counted += n
	}
}

func Merge[T cmp.Ordered](slice []T) []T {
	switch len(slice) {
	case 0, 1:
		return slice
	default:
		mid := len(slice) / 2
		arr1 := Merge(slice[:mid])
		arr2 := Merge(slice[mid+1:])
		return MergeSorted(arr1, arr2)
	}
}

// merges two sorted arrays into a single sorted one
func MergeSorted[T cmp.Ordered](A, B []T) []T {
	var (
		C      = make([]T, len(A)+len(B))
		filled = 0
	)

	for len(A) != 0 && len(B) != 0 {
		if A[0] < B[0] {
			C[filled] = A[0]
			A = A[1:]
		} else {
			C[filled] = B[0]
			B = B[1:]
		}
		filled++
	}

	for len(A) != 0 {
		C[filled] = A[0]
		A = A[1:]
		filled++
	}

	for len(B) != 0 {
		C[filled] = B[0]
		B = B[1:]
		filled++
	}

	return C
}

func Quick[T cmp.Ordered](slice []T) {
	switch len(slice) {
	case 0, 1:
		return
	default:
		Swap(slice, 0, rand.Intn(len(slice)))
		l := 0

		for r := 0; r < len(slice); r++ {
			if l != r && slice[r] < slice[0] {
				Swap(slice, r, l+1)
				l++
			}
		}

		Swap(slice, 0, l)

		Quick(slice[:l])
		Quick(slice[l+1:])
	}
}
