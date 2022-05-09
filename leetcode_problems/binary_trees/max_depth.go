package binary_trees

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func Max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}
