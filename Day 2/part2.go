package Day_2

func CalculateChecksum2(spreadsheet []string) int {
	sum := 0
	for _, stringRow := range spreadsheet {
		rowOfNumbers := SplitToNumbers(stringRow)
		divided, divider := GetDividers(rowOfNumbers)
		sum += divided / divider
	}

	return sum
}

func GetDividers(numbers []int) (int, int) {
	for _, divided := range numbers {
		for _, divider := range numbers {
			if divided != divider {
				if divided%divider == 0 {
					return divided, divider
				}
			}
		}
	}
	return 0, 0
}
