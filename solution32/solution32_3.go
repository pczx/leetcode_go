package solution32

func levelOrder_3(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	flag := false
	for len(queue) > 0 {
		level := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			root := queue[0]
			level = append(level, root.Val)
			queue = queue[1:]
			if root.Left != nil {
				queue = append(queue, root.Left)
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
			}
		}
		if flag {
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}
		res = append(res, level)
		flag = !flag
	}
	return res
}
