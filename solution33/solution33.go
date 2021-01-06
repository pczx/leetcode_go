package solution33

func verifyPostorder(postorder []int) bool {
	return judge(postorder, 0, len(postorder)-1)
}

func judge(postorder []int, start, end int) bool {
	if start >= end {
		return true
	}
	var i int
	for i = start; i < end; i++ {
		if postorder[i] > postorder[end] {
			break
		}
	}

	for j := i; j < end; j++ {
		if postorder[j] < postorder[end] {
			return false
		}
	}

	return judge(postorder, start, i-1) && judge(postorder, i, end-1)
}
