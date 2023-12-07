package set

// this is the most trivial set implementation possible
// the interface is very helpful, however

import "maps"

type Set[T comparable] map[T]struct{}

func (set Set[T]) Add(element T) {
	set[element] = struct{}{}
}

func (set Set[T]) In(element T) bool {
	_, in := set[element]
	return in
}

// intersection is in calling set
func (whom Set[T]) IntersectInPlace(with Set[T]) {
	for k := range whom {
		if _, ok := with[k]; !ok {
			delete(whom, k)
		}
	}
}

func (whom Set[T]) UnionInPlace(with Set[T]) {
	for k := range with {
		if _, ok := whom[k]; !ok {
			whom[k] = struct{}{}
		}
	}
}

func (whom Set[T]) DifferenceInPlace(with Set[T]) {
	for k := range whom {
		if _, ok := with[k]; ok {
			delete(whom, k)
		}
	}
}

func (set1 Set[T]) Union(set2 Set[T]) Set[T] {
	newMap := Set[T]{}
	maps.Copy(newMap, set1)
	newMap.UnionInPlace(set2)
	return newMap
}

func (set1 Set[T]) Intersect(set2 Set[T]) Set[T] {
	newMap := Set[T]{}
	for k := range set1 {
		if _, in := set2[k]; in {
			newMap.Add(k)
		}
	}
	return newMap
}
