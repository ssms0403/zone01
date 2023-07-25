package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		arg := os.Args[1]
		for _, c := range arg {
			repeat := 1
			if c >= 'a' && c <= 'z' {
				repeat += int(c) - 97
			} else if c >= 'A' && c <= 'Z' {
				repeat += int(c) - 65
			}
			for i := 0; i < repeat; i++ {
				z01.PrintRune(c)
			}
		}
		z01.PrintRune('\n')
	}
}
