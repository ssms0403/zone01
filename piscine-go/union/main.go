package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		allStrings := os.Args[1] + os.Args[2]
		unionString := []rune(nil)
		for _, aChar := range allStrings {
			isFound := false
			for _, uChar := range unionString {
				if aChar == uChar {
					isFound = true
					break
				}
			}
			if !isFound {
				unionString = append(unionString, aChar)
			}
		}
		for _, char := range unionString {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
