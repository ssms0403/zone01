package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		words := SplitWhiteSpaces(os.Args[1])
		fmt.Println(words)
		for i := 1; i < len(words); i++ {
			PrintStr(words[i])
			z01.PrintRune(' ')
		}
		PrintStr(words[0])
		z01.PrintRune('\n')
	}
}

func PrintStr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func SplitWhiteSpaces(s string) []string {
	arr := []string(nil)
	lastWhiteSpaceIndex := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if lastWhiteSpaceIndex != i {
				arr = append(arr, s[lastWhiteSpaceIndex:i])
			}
			lastWhiteSpaceIndex = i + 1
		} else if i == len(s)-1 {
			arr = append(arr, s[lastWhiteSpaceIndex:i+1])
			lastWhiteSpaceIndex = i + 1
		}
	}
	return arr
}
