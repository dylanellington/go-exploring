package general

 // This is the interface that allows for creating nested lists.
 // You should not implement it, or speculate about its implementation
 type NestedInteger struct {
 }

 // Return true if this NestedInteger holds a single integer, rather than a nested list.
 func (this NestedInteger) IsInteger() bool { return true }

 // Return the single integer that this NestedInteger holds, if it holds a single integer
 // The result is undefined if this NestedInteger holds a nested list
 // So before calling this method, you should have a check
 func (this NestedInteger) GetInteger() int { return -1 }

 // Set this NestedInteger to hold a single integer.
 func (n *NestedInteger) SetInteger(value int) {}

 // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 func (this *NestedInteger) Add(elem NestedInteger) {}

 // Return the nested list that this NestedInteger holds, if it holds a nested list
 // The list length is zero if this NestedInteger holds a single integer
 // You can access NestedInteger's List element directly if you want to modify it
 func (this NestedInteger) GetList() []*NestedInteger { return []*NestedInteger{}}

type NestedIterator struct {
	nestedList []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator {
		nestedList: nestedList,
	}
}

func (this *NestedIterator) Next() int {
	nestedIntegar := this.nestedList[0]
	this.nestedList = this.nestedList[1:]

	if nestedIntegar.IsInteger() {
		return nestedIntegar.GetInteger()
	} else {
		this.nestedList = append(nestedIntegar.GetList(), this.nestedList...)
		return this.Next()
	}
}

func (this *NestedIterator) HasNext() bool {
	for index := 0; index < len(this.nestedList); index++ {
		if this.nestedList[index].IsInteger() {
			return true
		}

		innerIterator := &NestedIterator {
			nestedList: this.nestedList[index].GetList(),
		}

		if innerIterator.HasNext() {
			return true
		}
	}

	return false
}
