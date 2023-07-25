package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		letter := os.Args[1]
		if len(letter) == 1 {
			letter := []rune(letter)[0]
			if letter >= 'a' && letter <= 'z' {
				for i := letter; i <= 'z'; i++ {
					z01.PrintRune(i)
				}
				for i := 'a'; i < letter; i++ {
					z01.PrintRune(i)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
