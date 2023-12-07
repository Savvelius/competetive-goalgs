package triev2

// this is a more optimised trie structure
// built as a DFA, not like a tree

import "unicode"

const SIZE = 26

type State = int

func toidx(ch byte) byte {
	return byte(unicode.ToLower(rune(ch))) - 97
}

type Trie struct {
	Q [][SIZE]State // i=state, j=transition. i=0 - invalid state. i=1 - start state

	accept map[State]struct{} // set of accept states
}

func New() Trie {
	return Trie{
		Q:      [][SIZE]State{{}, {}},
		accept: map[State]struct{}{},
	}
}

func (trie *Trie) NewState() int {
	trie.Q = append(trie.Q, [SIZE]State{})
	return len(trie.Q) - 1
}

// returns index of a state
func (trie *Trie) Add(s string) {
	if len(s) == 0 {
		return
	}

	state := 1
	for i := 0; i < len(s)-1; i++ {
		to := toidx(s[i])

		if trie.Q[state][to] == 0 {
			trie.Q[state][to] = trie.NewState()
		}

		state = trie.Q[state][to]
	}

	trie.accept[state] = struct{}{}
}

func (trie *Trie) Transition(from int, to string) State {
	if len(to) == 0 {
		return from
	}

	return trie.Transition(trie.Q[from][to[0]], to[1:])
}

func (trie *Trie) Contains(s string) bool {
	state := trie.Transition(1, s)
	_, contained := trie.accept[state]
	return contained
}
