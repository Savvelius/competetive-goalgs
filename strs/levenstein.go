package strs

/*
Moves:
1. remove an element
2. insert an element
3. change an element
Assert they cost the same
*/

func Distance(A, B string) int {
	var (
		N, M = len(A) + 1, len(B) + 1
		F    = make([][]int, N)
	)

	for i := range F {
		F[i] = make([]int, M)
		F[i][0] = i
	}
	for j := 1; j < M; j++ {
		F[0][j] = j
	}

	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			if A[i-1] == B[j-1] {
				F[i][j] = F[i-1][j-1]
			} else {
				F[i][j] = 1 + min(F[i-1][j], F[i][j-1], F[i-1][j-1])
			}
		}
	}

	return F[len(F)-1][len(F[0])-1]
}
