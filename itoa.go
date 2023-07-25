package piscine

func Itoa(n int) string {
	number := []rune(nil)
	isNegative := false
	if n < 0 {
		n = -n
		isNegative = true
	}

	for n != 0 {
		number = append(number, rune((n%10)+48))
		n /= 10
	}

	if isNegative {
		number = append(number, '-')
	}

	for i := 0; i < len(number)/2; i++ {
		number[i], number[len(number)-i-1] = number[len(number)-i-1], number[i]
	}

	return string(number)
}
