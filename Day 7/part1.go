package Day_7

import (
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
