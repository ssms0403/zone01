package main

import "github.com/01-edu/z01"

func main() {
	str := "zaybxcwdveuftgshriqjpkolnm\n"
	for _, char := range str {
		z01.PrintRune(char)
	}
}
