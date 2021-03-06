package day447

// NaiveIntegerExponentiation returns the result of x^y by calculating it
// naively. Runs in O(y) time due to y multiplications of x.
// y must be 0 or greater otherwise will panic.
func NaiveIntegerExponentiation(x, y uint) uint {
	result := uint(1)
	for i := uint(0); i < y; i++ {
		result *= x
	}

	return result
}

// IntegerExponentiationBySquaring calculates x^y.
// Runs in O(log y) time.
func IntegerExponentiationBySquaring(x, y uint) uint {
	if y == 0 {
		return 1
	}

	z := uint(1)

	for y > 1 {
		if y%2 == 0 {
			y /= 2
		} else {
			z *= x
			y = (y - 1) / 2
		}

		x *= x
	}

	return x * z
}
