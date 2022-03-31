package language_features

// SetBit sets the bit at a given position to 1.
// Position starts at 0 and shifts left.
func SetBit(integer *uint, position int) {
	*integer |= 1 << position
}

// ClearBit sets the bit at a given position to 0.
// Position starts at 0 and shifts left.
// Go docs on XOR: ^x is m^x with m = "all bits set to 1" for unsigned x
func ClearBit(integer *uint, position int) {
	*integer &= ^(1 << position)
}

// ToggleBit sets the bit at a given position to its compliment.
// Position starts at 0 and shifts left.
func ToggleBit(integer *uint, position int) {
	*integer ^= 1 << position
}