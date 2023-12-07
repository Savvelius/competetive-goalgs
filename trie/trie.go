package trie

// prefix trie in its classis form

const K = 26

// stores only lowercase words
type Vertex struct {
	to       [K]*Vertex
	distance int
	terminal bool
}

func getIdx(char byte) int {
	if char >= 'a' && char <= 'z' {
		return int(char - 97)
	} else if char >= 'A' && char <= 'Z' {
		return int(char - 65)
	} else {
		return -1
	}
}

// returns new vertex with terminal set to true
func NewVertex() *Vertex {
	return &Vertex{
		to:       [K]*Vertex{},
		terminal: false,
	}
}

func (root *Vertex) AddString(s string) {
	if len(s) == 0 {
		root.terminal = true
		return
	}

	idx := getIdx(s[0])
	if idx == -1 {
		panic("support only lower and upper case letters")
	}

	if root.to[idx] == nil {
		root.to[idx] = NewVertex()
		root.to[idx].distance = root.distance + 1
	}

	root.to[idx].AddString(s[1:])
}

func (root *Vertex) Contains(s string) bool {
	if len(s) == 0 {
		return root.terminal
	}

	idx := getIdx(s[0])
	if idx == -1 {
		panic("support only lower and upper case letters")
	}

	nextVertex := root.to[idx]
	if nextVertex == nil {
		return false
	}

	return nextVertex.Contains(s[1:])
}

// mark s end node as non-terminal
func (root *Vertex) Erase(s string) {
	if len(s) == 0 {
		root.terminal = false
		return
	}

	idx := getIdx(s[0])
	if idx == -1 {
		panic("support only lower and upper case letters")
	}

	nextVertex := root.to[idx]
	if nextVertex == nil {
		return
	}

	nextVertex.Erase(s[1:])
}
