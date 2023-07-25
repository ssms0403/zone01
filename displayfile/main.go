package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	length := len(args)

	if length > 1 {
		fmt.Println("Too many arguments")
	} else if length == 0 {
		fmt.Println("File name missing")
	} else if args[0] == "quest8.txt" {
		content, err := os.ReadFile(args[0])
		if err != nil {
			fmt.Println("File not found")
			return
		}
		fmt.Print(string(content))
	}
}
