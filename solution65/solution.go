package solution65

func add(a int, b int) int {
	carry := 0
	for b != 0 {
		carry = (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}
