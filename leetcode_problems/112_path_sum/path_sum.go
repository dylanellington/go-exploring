package _112_path_sum

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	newTargetSum := targetSum - root.Val

	if root.Left == nil && root.Right == nil && newTargetSum == 0 {
		return true
	}

	return hasPathSum(root.Left, newTargetSum) || hasPathSum(root.Right, newTargetSum)
}
