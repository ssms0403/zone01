package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		s := os.Args[1]
		runes := []rune(nil)
		lastWhiteSpaceIndex := 0
		for i := 0; i < len(s); i++ {
			if s[i] == ' ' && i != len(s)-1 {
				if lastWhiteSpaceIndex != i {
					runes = append(runes, []rune(s[lastWhiteSpaceIndex:i]+"   ")...)
				}
				lastWhiteSpaceIndex = i + 1
			} else if i == len(s)-1 {
				if s[i] != ' ' {
					runes = append(runes, []rune(s[lastWhiteSpaceIndex:i+1])...)
				} else {
					runes = append(runes, []rune(s[lastWhiteSpaceIndex:i])...)
				}
				lastWhiteSpaceIndex = i + 1
			}
		}
		fmt.Println(string(runes))
	}
}
