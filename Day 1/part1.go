package Day_1

func calculateCaptcha1(captcha string) int {
	length := len(captcha)
	sum := 0

	for i, thisRune := range captcha {
		nextIndex := i + 1
		if i == length-1 {
			nextIndex = 0
		}
		nextRune := []rune(captcha)[nextIndex]

		//converting ascii representation of a digit to actual int
		thisDigit := int(thisRune) - '0'
		nextDigit := int(nextRune) - '0'

		if thisDigit == nextDigit {
			sum += thisDigit
		}
	}

	return sum
}
