package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		if nb, err := strconv.Atoi(os.Args[1]); err == nil && nb >= 0 {
			nbHex := convertHex(nb)
			printStr(nbHex)
		} else {
			printStr("ERROR")
		}
	}
}

func convertHex(nb int) string {
	numbers := []rune(nil)
	base := "0123456789abcdef"
	for nb != 0 {
		numbers = append(numbers, rune(base[nb%16]))
		nb /= 16
	}
	for i := 0; i < len(numbers)/2; i++ {
		numbers[i], numbers[len(numbers)-i-1] = numbers[len(numbers)-i-1], numbers[i]
	}
	nbHex := string(numbers)
	return nbHex
}

func printStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
