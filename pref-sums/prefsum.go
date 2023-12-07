package prefsums

import "github.com/Savvelius/competetive-goalgs/intmath"

// Returns array of len=len(slice)+1 where i-th element is sum(slice[:i])
// If elements of slices are > 0 than result is in ascending order
func GetPrefixSums[I intmath.Number](slice []I) []I {
	sums := make([]I, len(slice)+1)

	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + slice[i-1]
	}

	return sums
}

func Get2DPrefixSums[I intmath.Signed](A [][]I) [][]I {
	N := len(A) + 1
	M := len(A[0]) + 1

	P := make([][]I, N)
	for i := range P {
		P[i] = make([]I, M)
	}

	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			P[i][j] = P[i-1][j] + P[i][j-1] - P[i-1][j-1] + A[i][j]
		}
	}

	return P
}

// reducable to two-sum
func HasSubArray[I intmath.Number](prefixes []I, sum I) bool {
	visited := make(map[I]struct{}, len(prefixes)-1)

	for i := 0; i < len(prefixes); i++ {
		pref := prefixes[i]
		_, isRightIdx := visited[pref+sum]
		_, isLeftIdx := visited[pref-sum]
		if isRightIdx || isLeftIdx {
			return true
		}
		visited[pref] = struct{}{}
	}

	return false
}

func ZeroLenSubArray[S intmath.Signed](p []S) int {
	m := map[S]int{}
	maxLen := 0

	for i := range p {
		prev, ok := m[p[i]]
		if ok {
			maxLen = max(maxLen, i-prev)
		}
		m[p[i]] = i
	}

	return maxLen
}

// some greedy alg. Might be true. Can't prove.
func SameSumSubArray[I intmath.Signed](p1, p2 []I) (int, int) {
	if len(p1) != len(p2) {
		panic("array lengths don't match")
	}

	l, r := 0, 1
	for r <= len(p1) && l <= len(p1) {
		sum1 := p1[r] - p1[l]
		sum2 := p2[r] - p2[l]
		if sum1 == sum2 {
			return l, r
		}

		if l >= r-1 {
			r++
		} else { // here we minimax
			if l == len(p1) {
				r++
			} else if r == len(p1) {
				l++
			} else {
				// try move r
				sum1 = p1[r] - p1[l+1]
				sum2 = p2[r] - p2[l+1]
				if sum1 == sum2 {
					return l, r
				}
				deltaL := intmath.Abs(sum1 - sum2)

				// try move r
				sum1 = p1[r+1] - p1[l]
				sum2 = p2[r+1] - p2[l]
				if sum1 == sum2 {
					return l, r
				}
				deltaR := intmath.Abs(sum1 - sum2)

				if deltaL < deltaR {
					l++
				} else {
					r++
				}
			}
		}
	}

	return -1, -1
}
