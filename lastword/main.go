package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg := []rune(os.Args[1])
		for i := len(arg) - 1; i >= 0; i-- {
			if arg[i] == ' ' {
				word := arg[i+1:]
				PrintStr(string(word))
				break
			}
		}
	}
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
