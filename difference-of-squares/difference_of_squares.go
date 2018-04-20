package diffsquares

func SumOfSquares(number int) int {
	squareOfSums := 0
	for i := number; i >= 0; i-- {
		squareOfSums += i * i
	}
	return squareOfSums
}

func SquareOfSums(number int) int {
	sumOfSquares := 0
	for i := number; i >= 0; i-- {
		sumOfSquares += i
	}
	return sumOfSquares * sumOfSquares
}

func Difference(number int) int {
	return SquareOfSums(number) - SumOfSquares(number)
}