package prefsums

import (
	"github.com/Savvelius/competetive-goalgs/intmath"
)

// reverse to prefix sum
func GetDiffArray[S intmath.Signed](slice []S) []S {
	n := len(slice) - 1

	d := make([]S, n)
	for i := range d {
		d[i] = slice[i+1] - slice[i]
	}

	return d
}

func Prepend[T any](slice []T, value T) []T {
	slice = append(slice, value)
	copy(slice[1:], slice)
	slice[0] = value
	return slice
}

// in real world there will be an array of queries
func AddRange[S intmath.Signed](arr []S, l, r int, value S) []S {
	arr = Prepend(arr, 0)
	ds := GetDiffArray(arr)

	ds[l] += value
	ds[r] -= value

	return GetPrefixSums(ds)[1:]
}

// adding arithmetic progression reduces to adding constant to range in diff array
