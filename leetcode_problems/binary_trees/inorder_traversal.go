package binary_trees

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int {}
	}

	if root.Left == nil && root.Right == nil {
		return []int { root.Val }
	}

	tree := [][]int {
		inorderTraversal(root.Left),
		[]int {root.Val},
		inorderTraversal(root.Right),
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
