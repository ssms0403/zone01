package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}
	area, err := strconv.Atoi(args[1])
	if err != nil || area < 0 {
		PrintStr("Error")
		return
	}
	var result int = int(float32(area*area) * float32(3.14))
	PrintStr(Itoa(result))
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

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
