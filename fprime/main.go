package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		number, _ := strconv.Atoi(os.Args[1])
		actualPrime := 2
		for number > 1 {
			if number%actualPrime == 0 {
				number /= actualPrime
				fmt.Print(actualPrime)
				if number != 1 {
					fmt.Print("*")
				} else {
					fmt.Println()
				}
			}
			if number%actualPrime != 0 {
				actualPrime = findNextPrime(actualPrime)
			}
		}
	}
}

func findNextPrime(nb int) int {
	for i := nb + 1; ; i++ {
		if isPrime(i) {
			return i
		}
	}
}

func isPrime(nb int) bool {
	for i := nb - 1; i > 1; i-- {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
