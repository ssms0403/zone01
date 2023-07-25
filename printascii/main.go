package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 2 || len(args[1]) != 1 {
		return
	}
	ascii := int(args[1][0])
	if ascii >= 65 && ascii <= 90 || ascii >= 97 && ascii <= 122 {
		printStr(itoa(ascii))
		return
	}
}

func printStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func itoa(nb int) string {
	number := []rune(nil)
	for {
		number = append(number, rune((nb%10)+48))
		nb /= 10
		if nb == 0 {
			break
		}
	}
	return string(number)
}
