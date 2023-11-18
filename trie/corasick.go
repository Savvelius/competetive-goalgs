package trie

// returns first occurence of any string in dict or -1 if not found. O(len(text)*VocabSize)
func AhoCorasick(dict []string, text string) int {
	root := NewVertex()
	for _, str := range dict {
		root.AddString(str)
	}

	cur := root

	for i, char := range []byte(text) {
		if cur.terminal {
			return i - cur.distance
		}

		idx := getIdx(char)
		if idx == -1 {
			cur = root
			continue
		}

		if cur.to[idx] == nil {
			if root.to[idx] == nil {
				cur = root
			} else {
				cur = root.to[idx]
			}
		} else {
			cur = cur.to[idx]
		}
	}

	return -1
}
