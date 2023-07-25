package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 2 {
		nbr, _ := strconv.Atoi(os.Args[1])
		printBits(nbr)
	}
}

func printBits(oct int) {
	arr := []int(nil)
	for bit := oct; bit > 0; bit >>= 1 {
		arr = append(arr, bit&1)
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 1 {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
	}
}
