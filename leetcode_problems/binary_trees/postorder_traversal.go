package binary_trees

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int {}
	}

	if root.Left == nil && root.Right == nil {
		return []int { root.Val }
	}

	tree := [][]int {
		postorderTraversal(root.Left),
		postorderTraversal(root.Right),
		[]int {root.Val},
	}

	treeSize := 0

	for _, node := range tree {
		treeSize += len(node)
	}

	result := make([]int, treeSize, treeSize)
	index := 0

	for _, node := range tree {
		index += copy(result[index:], node)
	}

	return result
}
