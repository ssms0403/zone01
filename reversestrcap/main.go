package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		arg := []rune(arg)
		for i, char := range arg {
			if i+1 == len(arg) || arg[i+1] == ' ' {
				if char >= 'a' && char <= 'z' {
					arg[i] = rune(char - 32)
				}
			} else if char >= 'A' && char <= 'Z' {
				arg[i] = rune(char + 32)
			}
		}
		fmt.Println(string(arg))
	}
}
