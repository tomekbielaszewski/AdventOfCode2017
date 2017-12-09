package Day_9

import (
	"container/list"
	"fmt"
)

func ComputeCollectedGarbage(code *list.List) int {
	collectedGarbage := 0
	for el := code.Front(); el != nil; el = el.Next() {
		line := el.Value.(string)
		chars := []rune(line)
		fmt.Println(string(chars))
		collectedGarbage += collectGarbage(chars)
	}
	return collectedGarbage
}

func collectGarbage(chars []rune) int {
	garbageCollected := 0
	isGarbageCollectingStarted := false
	omitNextChar := false

	for _, char := range chars {
		switch char {
		case '{':
			{
				if isGarbageCollectingStarted {
					if !omitNextChar {
						garbageCollected++
					}
				}
				omitNextChar = false
				break
			}
		case '}':
			{
				if isGarbageCollectingStarted {
					if !omitNextChar {
						garbageCollected++
					}
				}
				omitNextChar = false
				break
			}
		case '!':
			{
				if !omitNextChar {
					omitNextChar = true
				} else {
					omitNextChar = false
				}
				break
			}
		case ',':
			{
				if !omitNextChar && isGarbageCollectingStarted {
					garbageCollected++
				}
				omitNextChar = false
				break
			}
		case '<':
			{
				if !omitNextChar && isGarbageCollectingStarted {
					garbageCollected++
				}
				isGarbageCollectingStarted = true
				omitNextChar = false
				break
			}
		case '>':
			{
				if !omitNextChar {
					isGarbageCollectingStarted = false
				}
				omitNextChar = false
				break
			}
		default:
			{
				if !omitNextChar {
					garbageCollected++
				}
				omitNextChar = false
			}
		}
	}

	return garbageCollected
}
