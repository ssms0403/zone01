package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd(nb1, nb2 int) int {
	for nb1 != nb2 {
		if nb1 > nb2 {
			nb1 -= nb2
		} else {
			nb2 -= nb1
		}
	}
	return nb1
}

func main() {
	if len(os.Args) == 3 {
		nb1, _ := strconv.Atoi(os.Args[1])
		nb2, _ := strconv.Atoi(os.Args[2])
		fmt.Println(gcd(nb1, nb2))
	}
}
