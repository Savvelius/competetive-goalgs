package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := arr1
	arr1[2] = 1000
	fmt.Println(arr2)
}
