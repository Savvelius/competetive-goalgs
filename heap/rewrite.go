package heap

import (
	"math"
)

type Heap[T any] struct {
	Array  []T
	isLess func(T, T) bool
}

// Creates new heap. isLess should return true if its first argument is less than second.
// Panics on negative capacity.
func New[T any](isLess func(T, T) bool, capacity int) Heap[T] {
	if capacity < 0 {
		panic("capacity can't be negative")
	}
	return Heap[T]{
		Array:  make([]T, 0, capacity),
		isLess: isLess,
	}
}

// creates new heap from slice in O(n*log2(n))
func Heapify[T any](slice []T, isLess func(T, T) bool) Heap[T] {
	heap := New[T](isLess, len(slice))
	for _, t := range slice {
		heap.Insert(t)
	}
	return heap
}

// Returns slice sorted in ascending order from heap's values.
// Destroys heap
func (heap *Heap[T]) ToSlice() []T {
	slice := make([]T, 0, heap.Len())

	for heap.Len() != 0 {
		extracted := heap.ExtractMin()
		slice = append(slice, extracted)
	}

	return slice
}

func (heap *Heap[T]) Len() int {
	return len(heap.Array)
}

func (heap *Heap[T]) Insert(value T) {
	heap.Array = append(heap.Array, value)
	heap.siftUp(len(heap.Array) - 1)
}

// returns value of min element in O(log2(n)) time. Panics if heap is empty
func (heap *Heap[T]) ExtractMin() T {
	if len(heap.Array) == 0 {
		panic("heap is empty")
	}

	extracted := heap.Array[0]
	heap.Array[0] = heap.Array[len(heap.Array)-1]
	heap.Array = heap.Array[:len(heap.Array)-1]
	heap.siftDown(0)
	return extracted
}

// unlike extract min doesn't remove min element
func (heap *Heap[T]) GetMin() T {
	if len(heap.Array) == 0 {
		panic("heap is empty")
	}

	return heap.Array[0]
}

// extracts all values except min
func (heap *Heap[T]) ExtractNonMin() []T {
	extracted := heap.Array[1:]
	heap.Array = heap.Array[:1]

	return extracted
}

func (heap *Heap[T]) Clone() Heap[T] {
	return Heap[T]{
		Array:  append([]T{}, heap.Array...),
		isLess: heap.isLess,
	}
}

func (heap *Heap[T]) Delete(index int) {
	if index < 0 || index >= heap.Len() {
		panic("index out of bounds")
	}

	heap.Array[index] = heap.Array[heap.Len()-1]
	heap.Array = heap.Array[:heap.Len()-1]
	heap.siftDown(index)
}

func (heap *Heap[T]) siftUp(idx int) {
	if idx == 0 {
		return
	}
	if heap.greater(parentIdx(idx), idx) {
		heap.swap(idx, parentIdx(idx))
		heap.siftUp(parentIdx(idx))
	}
}

func (heap *Heap[T]) siftDown(idx int) {
	switch heap.numChildren(idx) {
	case 0:
		return
	case 1:
		if heap.hasLeftChild(idx) {
			if heap.greater(idx, leftChildIdx(idx)) {
				heap.swap(idx, leftChildIdx(idx))
			}
		} else if heap.hasRightChild(idx) {
			if heap.greater(idx, rightChildIdx(idx)) {
				heap.swap(idx, rightChildIdx(idx))
			}
		}
	case 2:
		if heap.greater(leftChildIdx(idx), rightChildIdx(idx)) {
			if heap.greater(idx, rightChildIdx(idx)) {
				heap.swap(idx, rightChildIdx(idx))
				heap.siftDown(rightChildIdx(idx))
				return
			}
		}

		if heap.greater(idx, leftChildIdx(idx)) {
			heap.swap(idx, leftChildIdx(idx))
			heap.siftDown(leftChildIdx(idx))
			return
		}

		if heap.greater(idx, rightChildIdx(idx)) {
			heap.swap(idx, rightChildIdx(idx))
			heap.siftDown(rightChildIdx(idx))
			return
		}
	default:
		panic("Can't happen")
	}
}

func parentIdx(i int) int {
	return int(math.Ceil(float64(i)/2) - 1)
}

func leftChildIdx(i int) int {
	return 2*i + 1
}

func rightChildIdx(i int) int {
	return 2*i + 2
}

func (heap *Heap[T]) swap(lhs, rhs int) {
	heap.Array[lhs], heap.Array[rhs] = heap.Array[rhs], heap.Array[lhs]
}

func (heap *Heap[T]) outOfBounds(idx int) bool {
	return idx >= len(heap.Array) || idx < 0
}

func (heap *Heap[T]) numChildren(idx int) int {
	num := 0
	if !heap.outOfBounds(leftChildIdx(idx)) {
		num++
	}
	if !heap.outOfBounds(rightChildIdx(idx)) {
		num++
	}
	return num
}

func (heap *Heap[T]) hasLeftChild(idx int) bool {
	return !heap.outOfBounds(leftChildIdx(idx))
}

func (heap *Heap[T]) hasRightChild(idx int) bool {
	return !heap.outOfBounds(rightChildIdx(idx))
}

// returns true if (value at lhs) > (value at rhs)
func (heap *Heap[T]) greater(lhs, rhs int) bool {
	return heap.isLess(heap.Array[rhs], heap.Array[lhs])
}
