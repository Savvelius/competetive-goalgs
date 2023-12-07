package unionfind

/*
OPS:
New :: []T -> UnionFind
Find :: T -> int
Union :: T -> T -> int
*/

// invariant - indexes[t] is always in parents
// sizes(not root) == 0
// sizes(root) == numChildren + 1
type UnionFind[T comparable] struct {
	indexes map[T]int
	parents []int
	sizes   []int
}

// caller guarantees uniqueness of elements in slice
func New[T comparable](slice []T) UnionFind[T] {
	indexes := map[T]int{}
	for _, el := range slice {
		if _, visited := indexes[el]; !visited {
			indexes[el] = len(indexes)
		}
	}

	uf := UnionFind[T]{
		indexes: indexes,
		parents: make([]int, len(indexes)),
		sizes:   make([]int, len(indexes)),
	}

	for i := 0; i < len(indexes); i++ {
		uf.parents[i] = i
		uf.sizes[i] = 1
	}

	return uf
}

// returns index of a root to which the needle belongs
func (haystack *UnionFind[T]) Find(needle T) int {
	needleIdx, ok := haystack.indexes[needle]
	if !ok {
		panic("object not in the map")
	}

	// guaranteed to terminate by structure's definition
	for {
		needleParentIdx := haystack.parents[needleIdx]
		if haystack.parents[needleParentIdx] == needleParentIdx {
			return needleParentIdx
		}
		needleIdx = needleParentIdx
	}
}

// returns index of root of merged tree
func (uf *UnionFind[T]) Union(first, second T) int {
	firstRoot := uf.Find(first)
	secondRoot := uf.Find(second)

	// tree of first element is less than that of second one
	if uf.sizes[firstRoot] < uf.sizes[secondRoot] {
		uf.parents[firstRoot] = secondRoot

		uf.sizes[secondRoot] += uf.sizes[firstRoot]
		uf.sizes[firstRoot] = 0

		return secondRoot
	}

	// tree of second element is less or equal to that of first one

	uf.parents[secondRoot] = firstRoot

	uf.sizes[firstRoot] += uf.sizes[secondRoot]
	uf.sizes[secondRoot] = 0

	return firstRoot
}
