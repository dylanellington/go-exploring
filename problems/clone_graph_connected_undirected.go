package problems

type Node struct {
	Val int
    Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	nodeMap := make(map[int]*Node)
	return cloneNode(node, nodeMap)
}

func cloneNode(node *Node, nodeMap map[int]*Node) *Node {
	newNode := Node {
		Val: node.Val,
		Neighbors: make([]*Node, 0),
	}

	nodeMap[newNode.Val] = &newNode

	for _, neighbor := range node.Neighbors {
		_, exists := nodeMap[neighbor.Val]

		if !exists {
			nodeMap[neighbor.Val] = cloneNode(neighbor, nodeMap)
		}

		newNode.Neighbors = append(newNode.Neighbors, nodeMap[neighbor.Val])
	}

	return &newNode
}
