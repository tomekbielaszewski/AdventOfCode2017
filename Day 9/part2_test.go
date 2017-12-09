package Day_9

import (
	"testing"
	"container/list"
	"fmt"
)

func TestPart2_Example1(t *testing.T) {
	input := list.New()
	input.PushFront("<>")

	result := ComputeCollectedGarbage(input)

	expected := 0
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example2(t *testing.T) {
	input := list.New()
	input.PushFront("<random characters>")

	result := ComputeCollectedGarbage(input)

	expected := 17
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example3(t *testing.T) {
	input := list.New()
	input.PushFront("<<<<>")

	result := ComputeCollectedGarbage(input)

	expected := 3
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example4(t *testing.T) {
	input := list.New()
	input.PushFront("<{!>}>")

	result := ComputeCollectedGarbage(input)

	expected := 2
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example5(t *testing.T) {
	input := list.New()
	input.PushFront("<!!>")

	result := ComputeCollectedGarbage(input)

	expected := 0
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example6(t *testing.T) {
	input := list.New()
	input.PushFront("<!!!>>")

	result := ComputeCollectedGarbage(input)

	expected := 0
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example7(t *testing.T) {
	input := list.New()
	input.PushFront("<{odi!a,<{i<a>")

	result := ComputeCollectedGarbage(input)

	expected := 10
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := ComputeCollectedGarbage(input)

	fmt.Printf("Part 2 solution is: %d \n\n", result)
}
