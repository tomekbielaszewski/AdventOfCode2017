package Day_3

import (
	"fmt"
	"testing"
)

func TestPart2_Example1(t *testing.T) {
	input := 1

	result := CalculateWierdSums(input)

	expected := 2
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example2(t *testing.T) {
	input := 330

	result := CalculateWierdSums(input)

	expected := 351
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example3(t *testing.T) {
	input := 59

	result := CalculateWierdSums(input)

	expected := 122
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example4(t *testing.T) {
	input := 805

	result := CalculateWierdSums(input)

	expected := 806
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Solution(t *testing.T) {
	input := 347991

	result := CalculateWierdSums(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
