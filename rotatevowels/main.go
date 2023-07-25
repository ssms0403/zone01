package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	args := ""
	vowels := []string(nil)
	result := ""
	if len(arguments) >= 1 {
		for i := 0; i < len(arguments); i++ {
			args += arguments[i]
			if i < len(arguments)-1 {
				args += " "
			}
		}
		for i := 0; i < len(args); i++ {
			if check(rune(args[i])) {
				vowels = append(vowels, string(args[i]))
			}
		}
		for i, j := 0, 0; i < len(args); i++ {
			if check(rune(args[i])) {
				result += vowels[len(vowels)-j-1]
				j++
			} else {
				result += string(args[i])
			}
		}
		for i := 0; i < len(result); i++ {
			z01.PrintRune(rune(result[i]))
		}
	}
	z01.PrintRune('\n')
}

func check(x rune) bool {
	if x == 'a' || x == 'A' || x == 'e' || x == 'E' || x == 'o' || x == 'O' || x == 'u' || x == 'U' || x == 'i' || x == 'I' {
		return true
	}
	return false
}
