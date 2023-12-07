package main

import "fmt"

func main() {
	FromSlice(1, 2, 3, 4, 5, 6).Map(func(e *int) {
		*e += 10
	}).Filter(func(e *int) bool {
		*e--
		return *e%2 == 0
	}).Map(func(e *int) {
		fmt.Println(*e)
	}).Do()
}
