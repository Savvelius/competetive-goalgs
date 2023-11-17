package strs

func Manacher(s string) []int {
	d1 := ManacherEven(s)
	d2 := ManacherOdd(s)
	for i := range d1 {
		d1[i] += d2[i]
	}
	return d1
}
func ManacherOdd(s string) []int {
	n := len(s)
	p := make([]int, n)

	l, r := -1, -1
	for i := 0; i < n; i++ {
		if i <= r {
			i_r_dist := r - i + 1
			i_sym := l + r - i
			p[i] = min(p[i_sym], i_r_dist)
		}

		for i+p[i] < n && i-p[i] >= 0 && s[i+p[i]] == s[i-p[i]] {
			p[i]++
		}

		if i+p[i]-1 > r {
			r = i + p[i] - 1
			l = i - p[i] + 1
		}
	}

	return p
}

// returns lengths of palindromes at their left center index
func ManacherEven(s string) []int {
	n := len(s)
	p := make([]int, n)
	p[n-1] = 0

	l, r := 0, 0
	for i := 0; i < n-1; i++ {
		// right center in bounds
		if i+1 <= r {
			i_r_dist := r - i
			i_sym := l + r - i - 1
			p[i] = min(p[i_sym], i_r_dist)
		}

		for i+p[i]+1 < n && i-p[i] >= 0 && s[i+p[i]+1] == s[i-p[i]] {
			p[i]++
		}

		if i+p[i] > r {
			r = i + p[i]
			l = i - p[i] + 1
		}
	}

	return p
}
