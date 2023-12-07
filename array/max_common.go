package array

import (
	"cmp"
)

// max common subsequence algorithm
func MCS[T cmp.Ordered](A, B []T) int {
	N := len(A)
	M := len(B)

	S := make([][]int, M+1)

	for i := range S {
		S[i] = make([]int, N+1)
	}

	for i := 1; i <= M; i++ {
		for j := 1; j <= N; j++ {
			if B[i-1] == A[j-1] {
				S[i][j] = S[i-1][j-1] + 1
			} else {
				S[i][j] = max(S[i-1][j], S[i][j-1])
			}
		}
	}

	return S[len(S)-1][len(S[0])-1]
}

// S[i] = max by (j < i : a[j] < a[i]) of (S[j])
// max common increasing subsequence algorithm
func MCIS[T cmp.Ordered](A []T) int {
	S := make([]int, len(A))

	for i := range S {
		m := 0
		for j := 0; j < i; j++ {
			if A[i] > A[j] && S[j] > m {
				m = S[j]
			}
		}
		S[i] = m + 1
	}

	return S[len(S)-1]
}
