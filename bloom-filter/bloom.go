package bloom

import (
	"log"

	"github.com/Savvelius/graph-algs/bitarray"
	"github.com/Savvelius/graph-algs/hashing"
)

type Filter struct {
	hashers []hashing.Hasher
	array   bitarray.BitArray
}

func New(numHashers uint) Filter {
	if numHashers == 0 {
		log.Fatalf("numHashers cannot be negative")
	}
	hashers := make([]hashing.Hasher, numHashers)
	return Filter{
		hashers: hashers,
		array:   bitarray.New(numHashers),
	}
}
