package day6

var occurrence map[string]int

func CalculateCyclesOfInfiniteLoop(memoryBlocks []int) int {
	set = make(map[string]bool)
	occurrence = make(map[string]int)
	steps := 0

	for !isRepeated(memoryBlocks) {
		saveOccurrenceAndStep(memoryBlocks, steps)
		processMemory(memoryBlocks)
		steps++
	}

	occurredOn := getOccurrence(memoryBlocks)

	return steps - occurredOn
}

func getOccurrence(memoryBlocks []int) int {
	asString := toString(memoryBlocks)
	occurredOn := occurrence[asString]
	return occurredOn
}

func saveOccurrenceAndStep(memoryBlocks []int, step int) {
	if !isRepeated(memoryBlocks) {
		asString := toString(memoryBlocks)
		occurrence[asString] = step
		set[asString] = true
	}
}