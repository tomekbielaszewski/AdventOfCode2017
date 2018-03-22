package day6

import (
	"fmt"
	"strings"
)

var set map[string]bool

func CalculateCycles(memoryBlocks []int) int {
	set = make(map[string]bool)
	steps := 0

	for !isRepeated(memoryBlocks) {
		saveOccurrence(memoryBlocks)
		processMemory(memoryBlocks)
		steps++
	}

	return steps
}

func processMemory(blocks []int) {
	value, index := getMax(blocks)
	blocks[index] = 0
	spread(value, index+1, blocks)
}

func spread(value int, index int, blocks []int) {
	for i := 0; i < value; i++ {
		if index >= len(blocks) {
			index = 0
		}

		blocks[index]++
		index++
	}
}

func getMax(numbers []int) (int, int) {
	max := 0
	index := 0

	for _index, _value := range numbers {
		if _value > max {
			index = _index
			max = _value
		}
	}

	return max, index
}

func isRepeated(sequence []int) bool {
	asString := toString(sequence)
	return set[asString]
}

func saveOccurrence(memoryBlocks []int) {
	asString := toString(memoryBlocks)
	set[asString] = true
}

func toString(ints []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ints)), ","), "[]")
}
