package differenceofsquares

func SquareOfSum(n int) int {
    // Gauss's Summation squared
    gauss := n*(n+1)/2
    return gauss * gauss
}

func SumOfSquares(n int) int { 
    // Square Pyramidal Number formula
    QPNformula := n*(n + 1)*(2* n + 1)/6
    return QPNformula
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
