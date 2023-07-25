package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		s := os.Args[1]
		str := os.Args[2]
		toFound := len(s)
		i := 0
		actual := s[i]
		for _, char := range str {
			if char == rune(actual) {
				i++
				toFound--
				if toFound == 0 {
					PrintStr(s)
					break
				}
				actual = s[i]
			}
		}
	}
}

func PrintStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
