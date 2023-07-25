package main

import "github.com/01-edu/z01"

func main() {
	str := "AZBYCXDWEVFUGTHSIRJQKPLOMN\n"
	for _, char := range str {
		z01.PrintRune(char)
	}
}
