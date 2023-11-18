package prefsums

import "github.com/Savvelius/competetive-goalgs/intmath"

func GetPrefixSums[I intmath.Number](slice []I) []I {
	sums := make([]I, len(slice)+1)

	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + slice[i-1]
	}

	return sums
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
