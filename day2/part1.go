package Day_2

func CalculateChecksum1(spreadsheet []string) int {
	sum := 0
	for _, stringRow := range spreadsheet {
		rowOfNumbers := SplitToNumbers(stringRow)
		min, max := GetEdgeValues(rowOfNumbers)
		sum += max - min
	}

	return sum
}

func GetEdgeValues(numbers []int) (int, int) {
	min := int(^uint(0) >> 1) //MAX_INTEGER
	max := 0
	for _, number := range numbers {
		if min > number {
			min = number
		}
		if max < number {
			max = number
		}
	}

	return min, max
}
