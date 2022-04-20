package problems

type Stack struct {
	items []rune
}

func NewStack() Stack {
	return Stack {
		items: make([]rune, 0),
	}
}

func (this *Stack) Push(item rune) {
	this.items = append(this.items, item)
}

func (this *Stack) Pop() {
	stackSize := len(this.items) - 1

	if stackSize < 0 {
		return
	}

	this.items = this.items[:len(this.items) - 1]
}

func (this *Stack) Peek() (rune, bool) {
	stackSize := len(this.items) - 1

	if stackSize < 0 {
		return 0, false
	}

	return this.items[len(this.items) - 1], true
}

func GetParenthesesPairingMap() map[rune]rune {
	return map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
}

func isValid(s string) bool {
	stack := NewStack()
	parenthesesPairingMap := GetParenthesesPairingMap()

	for index, character := range s {
		_, isOpeningCharacter := parenthesesPairingMap[character]

		if index == 0 || isOpeningCharacter {
			stack.Push(character)
			continue;
		}

		stackCharacter, _ := stack.Peek()
		closingCharacter, closingCharacterExists := parenthesesPairingMap[stackCharacter]

		if !closingCharacterExists || character != closingCharacter {
			return false
		}

		stack.Pop()
	}

	if _, stackCharacterExists := stack.Peek(); stackCharacterExists {
		return false
	}

	return true
} s
