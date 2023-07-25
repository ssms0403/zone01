package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		var alreadyPrint []rune
		str1 := os.Args[1]
		str2 := os.Args[2]
		for _, char := range str1 {
			for _, actual := range str2 {
				if char == actual && Index(string(alreadyPrint), string(char)) == -1 {
					z01.PrintRune(char)
					alreadyPrint = append(alreadyPrint, char)
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}
	for i := 0; i <= len(s)-len(toFind); i++ {
		subs := s[i : i+len(toFind)]
		if subs == toFind {
			return i
		}
	}
	return -1
}
