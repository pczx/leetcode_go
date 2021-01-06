package solution27

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Left != nil {
		mirrorTree(root.Left)
	}

	if root.Right != nil {
		mirrorTree(root.Right)
	}

	root.Left, root.Right = root.Right, root.Left

	return root
}
