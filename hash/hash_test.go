package hash

import (
	"testing"
)

func TestPolyHash(t *testing.T) {
	hasher := New(30, 100)
	a := "Hello"
	b := " friend"
	ab := a + b

	ha := hasher.Hash(a)
	hb := hasher.Hash(b)
	hab := hasher.Hash(ab)

	strHash := hasher.RemoveSuffix(hab, hb)
	testEqualHashes(t, strHash, ha)

	strHash = hasher.RemovePrefix(hab, ha)
	testEqualHashes(t, strHash, hb)

	strHash = hasher.ConcatHash(ha, hb)
	testEqualHashes(t, strHash, hab)
}

func TestSubStrHash(t *testing.T) {
	hasher := New(30, 100)
	str := "Hello friend"
	subStr := str[3:8]

	strHash := hasher.Hash(str)
	subStrHash := hasher.Hash(subStr)
	prefixes := hasher.GetPrefixHashes(str)

	if prefixes[len(prefixes)-1] != strHash.Hash {
		t.Errorf("expected hash of subStr[:len]=%d to be equal to strHash=%d",
			prefixes[len(prefixes)-1], strHash.Hash)
	}

	hash := hasher.HashSubString(prefixes, 3, 8)
	testEqualHashes(t, hash, subStrHash)
}

func testEqualHashes(t *testing.T, h1, h2 StringHash) {
	if h1.Hash != h2.Hash {
		t.Errorf("expected h1 hash to be %d, got %d", h1.Hash, h2.Hash)
	}
	if h1.StrLen != h2.StrLen {
		t.Errorf("expected h1 len to be %d, got %d", h1.StrLen, h2.StrLen)
	}
}
