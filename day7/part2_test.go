package Day_7

import (
	"fmt"
	"testing"
)

func TestPart2_Example(t *testing.T) {
	input := ReadFromFile("example.txt")

	result := GetCorrectWeight(input)

	expected := 60
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestDay7_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := GetCorrectWeight(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
