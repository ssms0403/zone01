package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		str := os.Args[1]
		number := 0
		counter := 0
		for _, char := range str {
			if char == 'C' {
				counter++
			} else if char == 'D' {
				counter--
			}
			if counter == 0 {
				number++
			}
		}
		fmt.Println(number)
	}
}
