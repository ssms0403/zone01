package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		latinWord := os.Args[1]
		if !isVowel(rune(latinWord[0])) {
			consBeforeVow := []rune(nil)
			for i, char := range latinWord {
				if !isVowel(char) {
					consBeforeVow = append(consBeforeVow, char)
				} else {
					latinWord = latinWord[i:] + string(consBeforeVow)
					break
				}
			}
			if string(consBeforeVow) == latinWord {
				printStr("No vowels")
				return
			}
		}
		printStr(latinWord + "ay")
	}
}

func isVowel(b rune) bool {
	return b == 'a' || b == 'i' || b == 'u' || b == 'o' || b == 'e' || b == 'A' || b == 'I' || b == 'U' || b == 'O' || b == 'E'
}

func printStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
