package Day_5

import (
	"testing"
	"fmt"
)

func TestPart1_Example1(t *testing.T) {
	input := [5]int{0, 3, 0, 1, -3}

	result := CalculateSteps(input[:])

	expected := 5
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := CalculateSteps(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
