package hashmap

import (
	"hash/fnv"
	"log"

	"github.com/Savvelius/competetive-goalgs/bintree"
	// TODO: write some balanced tree + maybe benchmark it vs std map - would be nice
)

func assert(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

const Size = 100

type Map[V any] struct {
	array  [Size]bintree.Node[uint32, V]
	hasher Hasher
}

func New[V any]() Map[V] {
	arr := [Size]bintree.Node[uint32, V]{}
	for i := range arr {
		arr[i] = bintree.Node[uint32, V]{}
	}
	return Map[V]{
		array:  arr,
		hasher: NewHasher(fnv.New32a()),
	}
}

func (mapa *Map[V]) getBucketAndHashAt(key any) (bintree.Node[uint32, V], uint32) {
	hash, err := mapa.hasher.Calculate(key)
	assert(err)
	idx := hash % Size
	return mapa.array[idx], hash
}

func (mapa *Map[V]) InsertOrUpdate(key any, value V) {
	bucket, hash := mapa.getBucketAndHashAt(key)
	bucket.InsertOrUpdate(hash, value)
}

func (mapa *Map[V]) Get(key any) *V {
	bucket, hash := mapa.getBucketAndHashAt(key)
	foundNode := bucket.Get(hash)
	if foundNode == nil {
		return nil
	}
	return &foundNode.Val
}

func (mapa *Map[V]) Contains(key any) bool {
	return mapa.Get(key) != nil
}
