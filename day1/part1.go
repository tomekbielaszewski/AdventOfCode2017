package Day_1

func CalculateCaptcha1(captcha string) int {
	length := len(captcha)
	sum := 0

	for i, thisRune := range captcha {
		nextIndex := i + 1
		if i == length-1 {
			nextIndex = 0
		}
		nextRune := []rune(captcha)[nextIndex]

		thisDigit := toInt(thisRune)
		nextDigit := toInt(nextRune)

		if thisDigit == nextDigit {
			sum += thisDigit
		}
	}

	return sum
}
