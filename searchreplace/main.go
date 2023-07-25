package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 4 {
		text := []rune(os.Args[1])
		search := []rune(os.Args[2])[0]
		replace := []rune(os.Args[3])[0]
		for i, char := range text {
			if char == search {
				text[i] = replace
			}
		}
		PrintStr(string(text))
	}
}

func PrintStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
