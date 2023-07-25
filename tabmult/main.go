package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		factor, _ := strconv.Atoi(os.Args[1])
		for i := 1; i <= 9; i++ {
			result := factor * i
			str := Itoa(i) + " x " + Itoa(factor) + " = " + Itoa(result)
			PrintStr(str)
		}
	}
}

func PrintStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func Itoa(nb int) string {
	runes := []rune(nil)
	for nb != 0 {
		char := (nb % 10) + 48
		runes = append(runes, rune(char))
		nb /= 10
	}
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
	return string(runes)
}
