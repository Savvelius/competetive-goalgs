package main

import (
	"fmt"

	prefsums "github.com/Savvelius/competetive-goalgs/pref-sums"
)

func main() {
	arr1 := []int{1, -2, 5, 8, 9, 123}
	pref1 := prefsums.GetPrefixSums(arr1)
	arr2 := []int{7, 4, 6, 2, 8, 111}
	pref2 := prefsums.GetPrefixSums(arr2)
	l, r := prefsums.SameSumSubArray(pref1, pref2)

	fmt.Println(l, r)
}
