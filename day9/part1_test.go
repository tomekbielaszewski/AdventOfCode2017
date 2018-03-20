package Day_9

import (
	"testing"
	"container/list"
	"fmt"
)

func TestPart1_Example1(t *testing.T) {
	input := list.New()
	input.PushFront("{}")

	result := ComputeScore(input)

	expected := 1
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example2(t *testing.T) {
	input := list.New()
	input.PushFront("{{{}}}")

	result := ComputeScore(input)

	expected := 6
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example3(t *testing.T) {
	input := list.New()
	input.PushFront("{{},{}}")

	result := ComputeScore(input)

	expected := 5
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example5(t *testing.T) {
	input := list.New()
	input.PushFront("{{{},{},{{}}}}")

	result := ComputeScore(input)

	expected := 16
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example6(t *testing.T) {
	input := list.New()
	input.PushFront("{<a>,<a>,<a>,<a>}")

	result := ComputeScore(input)

	expected := 1
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example7(t *testing.T) {
	input := list.New()
	input.PushFront("{{<ab>},{<ab>},{<ab>},{<ab>}}")

	result := ComputeScore(input)

	expected := 9
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example8(t *testing.T) {
	input := list.New()
	input.PushFront("{{<!!>},{<!!>},{<!!>},{<!!>}}")

	result := ComputeScore(input)

	expected := 9
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Example9(t *testing.T) {
	input := list.New()
	input.PushFront("{{<a!>},{<a!>},{<a!>},{<ab>}}")

	result := ComputeScore(input)

	expected := 3
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart1_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := ComputeScore(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
