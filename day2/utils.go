package Day_2

import (
	"strconv"
	"strings"
)

func SplitToNumbers(s string) []int {
	fields := strings.Fields(s)
	numbers := make([]int, len(fields))
	for i, field := range fields {
		numbers[i], _ = strconv.Atoi(field)
	}
	return numbers[:]
}
