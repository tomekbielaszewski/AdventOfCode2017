package Day_5

import (
	"testing"
	"fmt"
)

func TestPart2_Example1(t *testing.T) {
	input := [5]int{0, 3, 0, 1, -3}

	result := CalculateSteps2(input[:])

	expected := 10
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := CalculateSteps2(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
