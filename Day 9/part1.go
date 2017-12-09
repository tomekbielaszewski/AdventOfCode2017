package Day_9

import (
	"container/list"
	"fmt"
)

func ComputeScore(code *list.List) int {
	score := 0
	for el := code.Front(); el != nil; el = el.Next() {
		line := el.Value.(string)
		chars := []rune(line)
		fmt.Println(string(chars))
		score += computeScore(chars)
	}
	return score
}

func computeScore(chars []rune) int {
	score := 0
	charQueue := New()
	isGarbageCollectingStarted := false
	omitNextChar := false

	for _, char := range chars {
		switch char {
		case '{':
			{
				if !isGarbageCollectingStarted {
					charQueue.Push(char)
				}
				omitNextChar = false
				break
			}
		case '}':
			{
				if !isGarbageCollectingStarted {
					if charQueue.Len() > 0 {
						score += charQueue.Len()
						charQueue.Pop()
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
				omitNextChar = false
				break
			}
		case '<':
			{
				if !omitNextChar {
					isGarbageCollectingStarted = true
				}
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
				omitNextChar = false
			}
		}
	}

	return score
}
