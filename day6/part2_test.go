package day6

import (
	"testing"
	"fmt"
)

func TestPart2_Example1(t *testing.T) {
	input := [4]int{0, 2, 7, 0}

	result := CalculateCyclesOfInfiniteLoop(input[:])

	expected := 4
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := CalculateCyclesOfInfiniteLoop(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
