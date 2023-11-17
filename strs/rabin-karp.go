package strs

import "github.com/Savvelius/competetive-goalgs/hash"

func RabinKarp(haystack, needle string) []int {
	windowSize := len(needle)
	hasher := hash.New(65, windowSize)
	idxs := []int{}
	needleHash := hasher.Hash(needle).Hash
	prefixes := hasher.GetPrefixHashes(haystack)

	for i := 0; i+windowSize <= len(haystack); i++ {
		candidate := hasher.HashSubString(prefixes, i, i+windowSize).Hash

		if needleHash == candidate {
			idxs = append(idxs, i)
		}
	}

	return idxs
}
