package _144_binary_tree_preorder_traversal

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int {}
	}

	if root.Left == nil && root.Right == nil {
		return []int { root.Val }
	}

	tree := [][]int {
		[]int {root.Val},
		preorderTraversal(root.Left),
		preorderTraversal(root.Right),
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