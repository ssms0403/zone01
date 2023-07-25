package main

import (
	"fmt"
	"os"
)

func main() {
	numArg := len(os.Args)
	sum := 0
	if numArg == 2 {
		num := BasicAtoi(os.Args[1])
		for i := num; i > 1; i-- {
			if IsPrime(i) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
}

func BasicAtoi(s string) int {
	number := 0
	factor := 1
	for i := len(s) - 1; i >= 0; i-- {
		number += (int(s[i]) - 48) * factor
		factor *= 10
	}
	return number
}

func IsPrime(nb int) bool {
	for i := nb - 1; i > 1; i-- {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
