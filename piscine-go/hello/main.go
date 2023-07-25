package main

import "github.com/01-edu/z01"

func main() {
	str := "Hello World!"
	PrintStr(str)
}

func PrintStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
