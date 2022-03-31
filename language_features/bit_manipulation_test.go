package language_features

import (
	"testing"
)

func Test_Bit_Set(t *testing.T) {
	integer := uint(8)
	SetBit(&integer, 0)

	if integer != 9 {
		t.Errorf("Bit set on even number expected result to be 9, result was %v.", integer)
	}

	integer = 1
	SetBit(&integer, 2)

	if integer != 5 {
		t.Errorf("Bit set on odd number expected result to be 5, result was %v.", integer)
	}
}

func Test_Bit_Set_LargePosition(t *testing.T) {
	integer := uint(1)
	SetBit(&integer, 393738273)

	if integer != 1 {
		t.Errorf("Bit set on even number expected result to be 1, result was %v.", integer)
	}
}

func Test_Bit_Clear(t *testing.T) {
	integer := uint(10)
	ClearBit(&integer, 1)

	if integer != 8 {
		t.Errorf("Bit clear on even number expected result to be 8, result was %v.", integer)
	}

	integer = 1
	ClearBit(&integer, 0)

	if integer != 0 {
		t.Errorf("Bit clear on odd number expected result to be 0, result was %v.", integer)
	}
}

func Test_Bit_Toggle(t *testing.T) {
	integer := uint(10)
	ToggleBit(&integer, 1)

	if integer != 8 {
		t.Errorf("Bit toggle on even number (%v) expected result to be 8, result was %v.", 10, integer)
	}

	integer = 11
	ToggleBit(&integer, 0)

	if integer != 10 {
		t.Errorf("Bit toggle on odd number (%v) expected result to be 10, result was %v.", 11, integer)
	}
}

func Test_Bit_Check(t *testing.T) {
	integer := uint(4)
	isBitChecked := CheckBit(integer, 2)

	if isBitChecked != true {
		t.Errorf("Checked bit at postion %v for integar %v, expected result to be true. Result was %v.", 2, integer, isBitChecked)
	}
}

func Test_Bit_OnesCount(t *testing.T) {
	integer := uint(15)
	count := OnesCount(integer)

	if count != 4 {
		t.Errorf("Ones check of %v expected result to be 4, result was %v.", integer, count)
	}
}