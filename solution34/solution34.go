package solution34

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	paths := [][]int{}

	path := []int{}

	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, currSum int) {
		if root == nil {
			return
		}

		currSum += root.Val

		path = append(path, root.Val)

		if root.Left == nil && root.Right == nil && currSum == sum {
			temp := make([]int, len(path))
			copy(temp, path)
			paths = append(paths, temp)
		}

		if root.Left != nil {
			dfs(root.Left, currSum)
		}

		if root.Right != nil {
			dfs(root.Right, currSum)
		}

		path = path[:len(path)-1]
	}

	dfs(root, 0)

	return paths
}
