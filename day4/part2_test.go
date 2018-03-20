package Day_4

import (
	"testing"
	"fmt"
	"container/list"
)

func TestPart2_Example1(t *testing.T) {
	input := list.New()
	input.PushFront("abcde fghij")

	result := CountValidPassphrases2(input)

	expected := 1
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example2(t *testing.T) {
	input := list.New()
	input.PushFront("abcde xyz ecdab")

	result := CountValidPassphrases2(input)

	expected := 0
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example3(t *testing.T) {
	input := list.New()
	input.PushFront("a ab abc abd abf abj")

	result := CountValidPassphrases2(input)

	expected := 1
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example4(t *testing.T) {
	input := list.New()
	input.PushFront("iiii oiii ooii oooi oooo")

	result := CountValidPassphrases2(input)

	expected := 1
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Example5(t *testing.T) {
	input := list.New()
	input.PushFront("oiii ioii iioi iiio")

	result := CountValidPassphrases2(input)

	expected := 0
	if result != expected {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestPart2_Solution(t *testing.T) {
	input := ReadFromFile("input.txt")

	result := CountValidPassphrases2(input)

	fmt.Printf("Part 1 solution is: %d \n\n", result)
}
