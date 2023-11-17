package strs

// Returns indexes of starts of needle inside haystack.
// Asserts haystack doesn't contain byte(0)
func KMP(needle, haystack string) []int {
	prefixes := Z(needle + string(0) + haystack)
	answer := []int{}
	for i, pref := range prefixes {
		if pref == len(needle) {
			answer = append(answer, i-len(needle)-1)
		}
	}
	return answer
}

// O(N)
func Z(s string) []int {
	n := len(s)
	z := make([]int, n)

	// invariant: [l, r] is the right most prefix block
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			i_r_dist := r - i + 1
			i_sym_idx := i - l
			z[i] = min(z[i_sym_idx], i_r_dist)
		}

		for i+z[i] < n && s[i+z[i]] == s[z[i]] {
			z[i]++
		}

		// fix the invariant if broken
		if z[i] != 0 && i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}

	return z
}

// O(N^2)
func Zslow(s string) []int {
	n := len(s)
	z := make([]int, n)

	for i := 1; i < n; i++ {
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
	}

	return z
}

// TODO: PrefixToZ - algorithmica

func ZtoPrefix(z []int) []int {
	n := len(z)
	p := make([]int, n)

	for i := 1; i < n; i++ {
		if z[i] == 0 || z[i] == 1 {
			p[i] = max(z[i], p[i])
		} else {
			for j := 1; j <= z[i]; j++ {
				p[i+j-1] = max(j, p[i+j-1])
			}
		}
	}

	return p
}

// O(N)
func PrefixFn(s string) []int {
	n := len(s)
	p := make([]int, n)

	for i := 1; i < n; i++ {
		x := p[i-1]
	loop:
		if s[x] == s[i] {
			p[i] = x + 1
		} else if x != 0 {
			x = p[x-1]
			goto loop
		}
	}

	return p
}

// O(N^3)
func PrefixSlowFn(str string) []int {
	lengths := make([]int, len(str))

	for i := 1; i < len(str); i++ {
		for len := 1; len <= i; len++ {
			if str[:len] == str[i-len+1:i+1] {
				lengths[i] = len
			}
		}
	}

	return lengths
}
