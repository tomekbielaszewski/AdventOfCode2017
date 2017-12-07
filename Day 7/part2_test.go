package Day_7

import (
	"fmt"
	"testing"
)

func TestPart2_Example(t *testing.T) {
	input := ReadFromFile("example.txt")

	result, _ := GetCorrectWeight(input)

	expected := 60
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestDay7_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result, _ := GetCorrectWeight(input)

	fmt.Printf("Part 1 solution is: %s \n\n", result)
}
