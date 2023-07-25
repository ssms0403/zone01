package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	start := Atoi(args[1])
	files := args[2:]
	errors := false
	for i, fileName := range files {
		content, err := os.ReadFile(fileName)
		if err != nil {
			errors = true
			fmt.Printf("open %v: no such file or directory\n", fileName)
			continue
		}
		toPrint := string(content)
		_start := 0
		if len(toPrint)-start > 0 {
			_start = len(toPrint) - start
		}
		toPrint = toPrint[_start:]
		if len(files) == 1 {
			fmt.Printf("%v", toPrint)
		} else {
			if i >= 1 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %v <==\n%v", fileName, toPrint)
		}
	}
	if errors {
		os.Exit(1)
	}
}

func Atoi(s string) int {
	number := 0
	factor := 1
	for i := len(s) - 1; i >= 0; i-- {
		number += (int(s[i]) - 48) * factor
		factor = factor * 10
	}
	return number
}
