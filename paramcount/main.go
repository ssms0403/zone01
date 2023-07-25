package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	count := len(os.Args[1:])
	printStr(itoa(count))
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
