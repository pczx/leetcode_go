package solution44

import (
	"math"
)

func findNthDigit(n int) int {
	if n < 0 {
		return -1
	}
	digits := 1

	for {
		numbers := countOfInteger(digits)
		if n < numbers*digits {
			return digitAtIndex(n, digits)
		}

		n -= digits * numbers
		digits++
	}
}

func digitAtIndex(index, digits int) int {
	number := beginNumber(digits) + index/digits
	for i := digits - 1; i > index%digits; i-- {
		number /= 10
	}
	return number % 10
}

func beginNumber(digits int) int {
	if digits == 1 {
		return 0
	}
	return int(math.Pow(10.0, float64(digits-1)))
}

func countOfInteger(digits int) int {
	if digits == 1 {
		return 10
	}
	return 9 * int(math.Pow(10.0, float64(digits-1)))
}
