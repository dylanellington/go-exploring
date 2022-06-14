package _101_symmetric_tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	leftValues := GetNodeValues(root.Left, false)
	rightValues := GetNodeValues(root.Right, true)

	if len(leftValues) != len(rightValues) {
		return false
	}

	for index := 0; index < len(leftValues); index++ {
		if leftValues[index] != rightValues[index] {
			return false
		}
	}

	return true
}

func GetNodeValues(root *TreeNode, mirrorTraverse bool) []int {
	values := make([]int, 0)
	nodeStack := make([]*TreeNode, 0)
	nodeStack = append(nodeStack, root)

	for len(nodeStack) > 0 {
		node := nodeStack[len(nodeStack) - 1]
		nodeStack = nodeStack[0:len(nodeStack) - 1]

		if node != nil {
			values = append(values, node.Val)

			if mirrorTraverse {
				nodeStack = append(nodeStack, node.Right)
				nodeStack = append(nodeStack, node.Left)
			} else {
				nodeStack = append(nodeStack, node.Left)
				nodeStack = append(nodeStack, node.Right)
			}
		} else {
			values = append(values, -1)
		}
	}

	return values
}
