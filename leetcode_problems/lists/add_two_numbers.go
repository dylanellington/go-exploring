package lists

type ListNode struct {
	Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	currentNode := result
	carry := 0

	for count := 0; l1 != nil || l2 != nil; count++ {
		// value safety checks
		val1 := 0
		val2 := 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		// calculate sum and the carry for current nodes
		valueSum := val1 + val2 + carry

		if valueSum > 9 {
			carry = 1
			valueSum = valueSum - 10
		} else {
			carry = 0
		}

		// set node with calculated value
		if count == 0 {
			currentNode.Val = valueSum
		} else {
			currentNode.Next = new(ListNode)
			currentNode.Next.Val = valueSum

			currentNode = currentNode.Next
		}
	}

	// set final node if a carry still exists
	if carry > 0 {
		currentNode.Next = new(ListNode)
		currentNode.Next.Val = carry
	}

	return result
}
