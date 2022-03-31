package language_features

import (
	"testing"
)

func Test_Bit_Set(t *testing.T) {
	integer := uint(8)
	SetBit(&integer, 0)

	if integer != 9 {
		t.Errorf("Bit set on even number expected result to be 9, value was %v", integer)
	}

	integer = 1
	SetBit(&integer, 2)

	if integer != 5 {
		t.Errorf("Bit set on odd number expected result to be 5, value was %v", integer)
	}
}

func Test_Bit_Clear(t *testing.T) {
	integer := uint(10)
	ClearBit(&integer, 1)

	if integer != 8 {
		t.Errorf("Bit clear on even number expected result to be 8, value was %v", integer)
	}

	integer = 1
	ClearBit(&integer, 0)

	if integer != 0 {
		t.Errorf("Bit clear on odd number expected result to be 0, value was %v", integer)
	}
}

func Test_Bit_Toggle(t *testing.T) {
	integer := uint(10)
	ToggleBit(&integer, 1)

	if integer != 8 {
		t.Errorf("Bit toggle on even number expected result to be 8, value was %v", integer)
	}

	integer = 10
	ToggleBit(&integer, 0)

	if integer != 11 {
		t.Errorf("Bit toggle on odd number expected result to be 11, value was %v", integer)
	}
}