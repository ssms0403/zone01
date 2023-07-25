package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) >= 3 {
		for _, arg := range os.Args[1:] {
			for _, char := range arg {
				if (char < '0' || char > '9') && char != '-' {
					printStr("Error")
					return
				}
			}
		}
		min := os.Args[1]
		max := os.Args[len(os.Args)-1]
		str := min + " " + max
		printStr(str)
	}
}

func printStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
