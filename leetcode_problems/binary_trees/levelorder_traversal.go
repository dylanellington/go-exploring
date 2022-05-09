package binary_trees

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int {}
	}

	nodesQueue := make([]*TreeNode, 0)
	nodesQueue = append(nodesQueue, root)

	levels := [][]int {}
	nodesOnCurrentLevel := 1
	nodesOnNextLevel := 0

	for len(nodesQueue) > 0 {
		currentLevel := []int {}

		for index := 0; index < nodesOnCurrentLevel; index++ {
			if len(nodesQueue) == 0 {
				break
			}

			node := nodesQueue[0]

			if node != nil {
				currentLevel = append(currentLevel, node.Val)

				if node.Left != nil {
					nodesQueue = append(nodesQueue, node.Left)
					nodesOnNextLevel++
				}

				if node.Right != nil {
					nodesQueue = append(nodesQueue, node.Right)
					nodesOnNextLevel++
				}
			}

			nodesQueue = nodesQueue[1:]
		}

		levels = append(levels, currentLevel)
		nodesOnCurrentLevel = nodesOnNextLevel
		nodesOnNextLevel = 0
	}

	return levels
}
