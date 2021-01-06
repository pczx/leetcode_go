package solution31

func validateStackSequences(pushed []int, popped []int) bool {
	res := []int{}
	i := 0
	for _, v := range pushed {
		res = append(res, v)
		for len(res) > 0 && popped[i] == res[len(res)-1] {
			res = res[:len(res)-1]
			i++
		}

	}
	return len(res) == 0
}
