package Day_7

import (
	"container/list"
	"strings"
)

func GetBottomNodeName(input *list.List) (string, error) {
	nodes := make(map[string][]string)
	nodesCopy := make(map[string][]string)
	programNames := list.New()

	for e := input.Front(); e != nil; e = e.Next() {
		row := strings.Replace(e.Value.(string), ",", "", -1)
		rowElements := strings.Fields(row)
		nodes[rowElements[0]] = rowElements
		nodesCopy[rowElements[0]] = rowElements
		programNames.PushFront(rowElements[0])
	}

	for e := programNames.Front(); e != nil; e = e.Next() {
		rowElements := nodesCopy[e.Value.(string)]

		if len(rowElements) > 2 {
			for i := 3; i < len(rowElements); i++ {
				childElem := rowElements[i]
				delete(nodes, childElem)
			}
		}
	}

	for key := range nodes {
		return key, nil
	}

	var err error
	return "", err
}
