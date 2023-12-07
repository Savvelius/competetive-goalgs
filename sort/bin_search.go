package sort

type Num interface {
	int | uint | int64 | uint64 | int32 | uint32 | int16 | uint16 | int8 | uint8
}

// Requires function to be ascending on [l, r).
// Put greater-equals if function is descending.
func Rsearch[I Num](l, r I, less_eq func(I) bool) I {
	// invariant: f(l) <= needle, f(r) > needle
	// loop ends when r = l+1 -> as element at l <= target and
	// element at r is the first that > needle -> l is position of last
	// element <= needle
	for l < r-1 {
		m := (l + r) / 2
		if less_eq(m) {
			l = m
		} else {
			r = m
		}
	}
	return l
}

// Requires function to be ascending on (l, r].
// Put less-equals if function is descending.
func Lsearch[I Num](l, r I, greater_eq func(I) bool) I {
	// invariant: f(l) < needle, f(r) >= needle
	// loop ends when r = l+1 => as element at l is strictly less =>
	// element at r is the first that >= needle
	for l < r-1 {
		m := (l + r) / 2
		if greater_eq(m) {
			r = m
		} else {
			l = m
		}
	}
	return r
}
