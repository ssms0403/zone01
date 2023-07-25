package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) >= 2 {
		args := os.Args[1:]
		for i := len(args) - 1; i >= 0; i-- {
			printStr(args[i])
		}
	}
	z01.PrintRune('\n')
}

func printStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune(' ')
}
