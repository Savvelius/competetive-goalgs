package main

import (
	"fmt"

	"github.com/Savvelius/competetive-goalgs/strs"
)

func main() {
	s := "choose"
	t :=
		"choose life. choose a job. choose a career. choose a family. choose a fu..."

	idxs := strs.RabinKarp(t, s)

	for _, idx := range idxs {
		fmt.Println(t[idx : idx+len(s)])
	}
}
