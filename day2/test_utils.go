package Day_2

import "testing"

func AssertArrayIndexToBe(array []int, index int, expected int, t *testing.T) {
	if array[index] != expected {
		t.Error("Expected 1 on 0 index but got ", array[index])
	}
}
