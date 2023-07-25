package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		num, err := strconv.Atoi(os.Args[1])
		if err != nil {
			return
		}
		if num%2 == 0 {
			ping(num)
		} else {
			pong(num)
		}
	}
}

func ping(n int) {
	if n == 0 {
		return
	}
	printStr(itoa(n) + " ping")
	pong(n - 1)
}

func pong(n int) {
	if n == 0 {
		return
	}
	printStr(itoa(n) + " pong")
	ping(n - 1)
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
