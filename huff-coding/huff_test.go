package huff

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	inputs := []Input{
		{freq: 30, label: "A"},
		{freq: 51, label: "B"},
		{freq: 5, label: "C"},
		{freq: 146, label: "D"},
		{freq: 676, label: "E"},
		{freq: 166, label: "F"},
		{freq: 60, label: "G"},
		{freq: 60000, label: "H"},
	}
	root := Build(inputs)
	root.Pprint()
	fmt.Printf("Actual information: %f\n", root.LanguageInformation())
	fmt.Printf("Entropy: %f\n", root.LanguageEntropy())
}
