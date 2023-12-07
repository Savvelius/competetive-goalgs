package optimaltree

import (
	"cmp"
)

type Node[K cmp.Ordered] struct {
	key  K
	freq uint
}

func SumFreqs[T cmp.Ordered](nodes []Node[T]) uint {
	sum := uint(0)
	for _, node := range nodes {
		sum += node.freq
	}
	return sum
}

// assertion: nodes are ordered by keys in ascending order
// Note: memory consumption can be easily halfed
func Generate[K cmp.Ordered](nodes []Node[K]) uint {
	W := make([][]uint, len(nodes)+1)
	for i := range W {
		W[i] = make([]uint, len(nodes)+1)
	}

	// i = j + 1 - base case
	for i := 1; i < len(W); i++ {
		W[i][i-1] = 0
	}

	for s := 0; s < len(nodes); s++ {
		for i := 0; i < len(nodes)-s; i++ {
			minW := uint(0)
			for r := i; r < i+s; r++ {
				rW := W[i][r-1] + W[r+1][i]
				minW = min(minW, rW)
			}
			W[i][i+s] = SumFreqs(nodes[i:i+s]) + minW
		}
	}

	return W[1][len(W[0])-1]
}
