package language_features

import (
	"math/bits"
)

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

// CheckBit checks the bit at the given position and returns true if it is a 1.
// Position starts at 0 and shifts left.
func CheckBit(integer uint, position int) bool {
	return (integer & (1 << position)) > 0
}

// OnesCount returns the count of 1's in the given integer.
func OnesCount(integer uint) int {
	return bits.OnesCount(integer)
}