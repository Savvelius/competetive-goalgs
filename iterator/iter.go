package main

type Iterator[T any, Elem any, I any] struct {
	data T
	cur  I
	next func(T, *I) *Elem
}

func FromSlice[T any](slice ...T) *Iterator[[]T, T, int] {
	return &Iterator[[]T, T, int]{
		data: slice,
		next: func(t []T, i *int) *T {
			if *i >= len(t) {
				return nil
			}
			*i++
			return &t[*i-1]
		},
	}
}

func (iter *Iterator[T, Elem, I]) Next() *Elem {
	return iter.next(iter.data, &iter.cur)
}

func (iter *Iterator[T, Elem, I]) Map(f func(*Elem)) *Iterator[T, Elem, I] {
	return &Iterator[T, Elem, I]{
		data: iter.data,
		cur:  iter.cur,
		next: func(t T, i *I) *Elem {
			n := iter.next(t, i)
			if n != nil {
				f(n)
			}
			return n
		},
	}
}

func (iter *Iterator[T, Elem, I]) Filter(f func(*Elem) bool) *Iterator[T, Elem, I] {
	return &Iterator[T, Elem, I]{
		data: iter.data,
		cur:  iter.cur,
		next: func(t T, i *I) *Elem {
		Start:
			n := iter.next(t, i)
			if n == nil {
				return nil
			}
			if f(n) {
				return n
			}
			goto Start
		},
	}
}

func (iter *Iterator[T, Elem, I]) Do() {
	for v := iter.Next(); v != nil; v = iter.Next() {
	}
}
