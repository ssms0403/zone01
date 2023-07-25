package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		number, _ := strconv.Atoi(os.Args[1])
		if number == 1 {
			fmt.Println("false")
			return
		}
		for number != 1 {
			if number%2 != 0 {
				fmt.Println("false")
				return
			}
			number /= 2
		}
		fmt.Println("true")
	}
}
