package Day_7

import (
	"os"
	"bufio"
	"container/list"
	"strings"
)

func GetBottomNodeName(input *list.List) (string, error) {
	programs := make(map[string][]string)
	programsCopy := make(map[string][]string)
	programNames := list.New()

	for e := input.Front(); e != nil; e = e.Next() {
		row := strings.Replace(e.Value.(string), ",", "", -1)
		rowElements := strings.Fields(row)
		programs[rowElements[0]] = rowElements
		programsCopy[rowElements[0]] = rowElements
		programNames.PushFront(rowElements[0])
	}

	for e := programNames.Front(); e != nil; e = e.Next() {
		rowElements := programsCopy[e.Value.(string)]

		if len(rowElements) > 2 {
			for i := 3; i < len(rowElements); i++ {
				childElem := rowElements[i]
				delete(programs, childElem)
			}
		}
	}

	for key := range programs {
		return key, nil
	}

	var err error
	return "", err
}

func ReadFromFile(filename string) *list.List {
	stringLines := list.New()

	file, err := os.Open(filename)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringLines.PushFront(scanner.Text())
	}

	return stringLines
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
