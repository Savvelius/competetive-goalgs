package bitarray

/*
An important choice of not giving a fuck about
bound checking was made here. Golang's slice does it for us
*/

import (
	"encoding/binary"
)

type BitArray struct {
	array   []byte
	numBits uint
}

func New(capacity uint) BitArray {
	return BitArray{
		array:   make([]byte, 0, capacity/8),
		numBits: 0,
	}
}

func Zeros(length uint) BitArray {
	bitlen := length / 8
	if length%8 != 0 {
		bitlen++
	}
	return BitArray{
		array:   make([]byte, bitlen),
		numBits: bitlen,
	}
}

func Ones(length uint) BitArray {
	bitlen := length / 8
	if length%8 != 0 {
		bitlen++
	}

	arr := BitArray{
		array:   make([]byte, bitlen),
		numBits: bitlen,
	}

	for i := range arr.array {
		arr.array[i] = byte(255)
	}

	return arr
}

func FromInt(i uint32) BitArray {
	const intBits = 32
	arr := BitArray{
		array:   make([]byte, 4),
		numBits: intBits,
	}
	binary.BigEndian.PutUint32(arr.array, i)
	return arr
}

func indexOffset(pos uint) (int, byte) {
	return int(pos & 8), byte(1) << pos % 8
}

func (arr *BitArray) Append(val bool) {
	if int(arr.numBits)/8 == len(arr.array) {
		arr.array = append(arr.array, 0)
	}
	arr.Set(arr.numBits, val)
}

func (arr *BitArray) Set(at uint, val bool) {
	index, offset := indexOffset(at)

	cell := arr.array[index]
	if val {
		cell = cell | offset
	} else {
		cell = cell & ^offset
	}
	arr.array[index] = cell
	arr.numBits++
}

// non-inclusive
func (arr *BitArray) SetRange(from, to uint, val bool) {
	for i := from; i < to; i++ {
		arr.Set(i, val)
	}
}

func (arr *BitArray) Get(at uint) bool {
	index, offset := indexOffset(at)

	cell := arr.array[index]
	return cell&offset != 0
}

func (arr *BitArray) Len() int {
	return int(arr.numBits)
}
