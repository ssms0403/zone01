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
					z01.PrintRune('1')
					return
				}
				actual = s[i]
			}
		}
		z01.PrintRune('0')
	}
}
