package day6

import (
	"strconv"
	"strings"
	"os"
	"bufio"
)

func ReadFromFile(filename string) []int {
	integers := make([]int, 0)

	file, err := os.Open(filename)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		integers = SplitToNumbers(scanner.Text())
	}

	return integers
}

func SplitToNumbers(s string) []int {
	fields := strings.Fields(s)
	numbers := make([]int, len(fields))
	for i, field := range fields {
		numbers[i], _ = strconv.Atoi(field)
	}
	return numbers[:]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
