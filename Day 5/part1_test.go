package Day_5

import (
	"testing"
	"fmt"
)

func TestPart1_Example1(t *testing.T) {
	input := make([]int, 0)

	result := CalculateSteps(input)

	expected := 0
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := CalculateSteps(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
