package Day_1

func CalculateCaptcha2(captcha string) int {
	length := len(captcha)
	stepsForward := length/2
	sum := 0

	for i, thisRune := range captcha {
		nextIndex := GetCircularIndex(i + stepsForward, length)
		nextRune := []rune(captcha)[nextIndex]

		thisDigit := toInt(thisRune)
		nextDigit := toInt(nextRune)

		if thisDigit == nextDigit {
			sum += thisDigit
		}
	}

	return sum
}

func GetCircularIndex(bigIndex int, length int) int {
	return bigIndex % length
}