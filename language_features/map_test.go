package language_features

import (
	"testing"
)

func Test_Map_Make(t *testing.T) {
	hashmap := make(map[string]bool)
	hashmap["x"] = true
	hashmap["y"] = true

	value, exists := hashmap["x"]

	if !exists || value == false {
		t.Errorf("Hashmap literal did not return the expected value.")
	}
}

func Test_Map_Literal(t *testing.T) {
	hashmap := map[string]bool {
		"x": true,
		"y": true,
	}

	value, exists := hashmap["x"]

	if !exists || value == false {
		t.Errorf("Hashmap literal did not return the expected value.")
	}
}