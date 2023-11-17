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
	sum := uint64(0)

	for i := range s {
		sum = (sum + uint64(s[i])*hasher.NthPower(i)) % hasher.mod
	}
	return StringHash{
		StrLen: len(s),
		Hash:   sum,
	}
}

func (hasher *Hasher) ConcatHash(lhs, rhs StringHash) StringHash {
	hash := lhs.Hash + hasher.NthPower(lhs.StrLen)*rhs.Hash
	return StringHash{
		StrLen: lhs.StrLen + rhs.StrLen,
		Hash:   hash,
	}
}

func (hasher *Hasher) RemovePrefix(str, prefix StringHash) StringHash {
	hash := (str.Hash - prefix.Hash) / hasher.NthPower(prefix.StrLen)
	return StringHash{
		StrLen: str.StrLen - prefix.StrLen,
		Hash:   hash,
	}
}

func (hasher *Hasher) RemoveSuffix(str, suffix StringHash) StringHash {
	hash := str.Hash - suffix.Hash*hasher.NthPower(suffix.StrLen)
	return StringHash{
		StrLen: str.StrLen - suffix.StrLen,
		Hash:   hash,
	}
}

// their lengths are determined by indexes
func (hasher *Hasher) GetPrefixHashes(s string) []uint64 {
	n := len(s)
	hashes := make([]uint64, n+1)

	for i := 1; i < len(hashes); i++ {
		hashes[i] = (hashes[i-1] + hasher.NthPower(i)*uint64(s[i-1])) % hasher.mod
	}
	return hashes
}

// WARNING: this will not be equal to Hash(str[l:r]). r not incusive
func (hasher *Hasher) HashSubString(prefixHashes []uint64, l, r int) uint64 {
	return intmath.Pow(hasher.k, l) * (prefixHashes[r] - prefixHashes[l])
}
