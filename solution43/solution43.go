package solution43

import (
	"math"
	"strconv"
)

func countDigitOne(n int) int {
	return numberOf1(n)
}

func numberOf1(n int) int {
	if n <= 0 {
		return 0
	}

	s := strconv.Itoa(n)

	first := int(s[0] - '0')
	pow := int(math.Pow(10.0, float64(len(s)-1)))
	last := n - first*pow

	numberOfFirstDigit := 0
	if first == 1 {
		numberOfFirstDigit = last + 1
	} else {
		numberOfFirstDigit = pow
	}

	numberOfOtherDigit := first * numberOf1(pow-1)

	numRecrusive := numberOf1(last)

	return numberOfFirstDigit + numberOfOtherDigit + numRecrusive
}
