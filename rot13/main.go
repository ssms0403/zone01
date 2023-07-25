package main

import (
	"fmt"
	"os"
)

func Rot13(s string) string {
	runes := []rune(s)
	for i, c := range runes {
		if c >= 'a' && c <= 'z' {
			runes[i] = ((c - 97 + 13) % 26) + 97
		} else if c >= 'A' && c <= 'Z' {
			runes[i] = ((c - 65 + 13) % 26) + 65
		}
	}
	return string(runes)
}

func main() {
	if len(os.Args) == 2 {
		word := os.Args[1]
		r := Rot13((word))
		fmt.Println(string(r))
	}
}
