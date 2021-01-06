package leetcode337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	f := make(map[*TreeNode]int)
	g := make(map[*TreeNode]int)

	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		f[root] = root.Val + g[root.Left] + g[root.Right]
		g[root] = max(f[root.Left], g[root.Left]) + max(f[root.Right], g[root.Right])
	}

	dfs(root)
	return max(f[root], g[root])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
