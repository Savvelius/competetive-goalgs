package hash

import (
	"math"

	"github.com/Savvelius/competetive-goalgs/intmath"
)

type Hasher struct {
	k      uint64
	mod    uint64
	Powers []uint64
}

// naive O(N^2) algorithm.
func maxCoPrime(k uint64) uint64 {
	limit := uint64(math.MaxUint64)
	for !intmath.CoPrimes(k, limit) {
		limit--
	}
	return limit
}

func New(base uint64, powersSize int) Hasher {
	if powersSize <= 0 {
		panic("can't create not positive len powers for a hasher")
	}
	mod := maxCoPrime(base)

	hasher := Hasher{
		k:      base,
		mod:    mod,
		Powers: make([]uint64, powersSize),
	}

	hasher.Powers[0] = 1
	for i := 1; i < int(powersSize); i++ {
		hasher.Powers[i] = (hasher.Powers[i-1] * base) % mod
	}

	return hasher
}

func (hasher *Hasher) NthPower(n int) uint64 {
	if n < 0 {
		panic("support only positive powers")
	}
	for n >= len(hasher.Powers) {
		lastPow := hasher.Powers[len(hasher.Powers)-1]
		hasher.Powers = append(hasher.Powers, (lastPow*hasher.k)%hasher.mod)
	}
	return hasher.Powers[n]
}

type StringHash struct {
	// can store 4 additional bytes here!
	StrLen int
	Hash   uint64
}

// computes (s0​+s1​*k+s2​*k^2+…+sn*​k^n)
func (hasher *Hasher) Hash(s string) StringHash {
	n := len(s)
	sum := uint64(0)

	for i := n - 1; i >= 0; i-- {
		sum = (sum + uint64(s[n-1-i])*hasher.NthPower(i)) % hasher.mod
	}

	return StringHash{
		StrLen: n,
		Hash:   sum,
	}
}

func (hasher *Hasher) ConcatHash(lhs, rhs StringHash) StringHash {
	hash := lhs.Hash*hasher.NthPower(rhs.StrLen) + rhs.Hash
	return StringHash{
		StrLen: lhs.StrLen + rhs.StrLen,
		Hash:   hash,
	}
}

func (hasher *Hasher) RemovePrefix(str, prefix StringHash) StringHash {
	newLen := str.StrLen - prefix.StrLen
	hash := str.Hash - prefix.Hash*hasher.NthPower(newLen)

	return StringHash{
		StrLen: newLen,
		Hash:   hash,
	}
}

func (hasher *Hasher) RemoveSuffix(str, suffix StringHash) StringHash {
	newLen := str.StrLen - suffix.StrLen
	hash := (str.Hash - suffix.Hash) / hasher.NthPower(suffix.StrLen)

	return StringHash{
		StrLen: newLen,
		Hash:   hash,
	}
}

// Their lengths are determined by indexes.
// i-th prefix is Hash(s[:i]). Total of len(s)+1 hashes
func (hasher *Hasher) GetPrefixHashes(s string) []uint64 {
	hashes := make([]uint64, len(s)+1)

	for i := 0; i < len(s); i++ {
		hashes[i+1] = (hashes[i]*hasher.k + uint64(s[i])) % hasher.mod
	}
	return hashes
}

func (hasher *Hasher) HashSubString(prefixHashes []uint64, l, r int) StringHash {
	newLen := r - l
	hash := prefixHashes[r] - prefixHashes[l]*hasher.NthPower(newLen)

	return StringHash{
		StrLen: newLen,
		Hash:   hash,
	}
}
