package Day_4

import (
	"container/list"
	"strings"
)

func CountValidPassphrases(input *list.List) int {
	invalid := 0
	for e := input.Front(); e != nil; e = e.Next() {
		passphrase := e.Value.(string)
		words := strings.Fields(passphrase)
		isPresent := make(map[string]bool)

		for _, word := range words {
			if isPresent[word] {
				invalid++
				break
			}
			isPresent[word] = true
		}
	}
	return input.Len() - invalid
}
