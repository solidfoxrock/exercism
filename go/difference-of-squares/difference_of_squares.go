package diffsquares

// SquareOfSums calculates square of the sums
func SquareOfSums(n int) int {
	return n * (n + 1) * n * (n + 1) / 4
}

// SquareOfSums calculates sum of the squares
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates difference between SquareOfSums and SumOfSquares
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
