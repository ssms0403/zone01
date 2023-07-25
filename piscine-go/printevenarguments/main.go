package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 2 {
		z01.PrintRune('\n')
		return
	}
	for i, arg := range os.Args[1:] {
		if i%2 != 0 {
			printStr(arg)
		}
	}
}

func printStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
