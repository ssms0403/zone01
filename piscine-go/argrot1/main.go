package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		args := os.Args[1:]
		for i := 1; i < len(args); i++ {
			fmt.Print(args[i] + " ")
		}
		fmt.Print(args[0])
	}
}
