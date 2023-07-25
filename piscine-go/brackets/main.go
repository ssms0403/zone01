package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		bracketStack := []rune(nil)
		str := os.Args[1]
		for _, char := range str {
			if char == '(' {
				bracketStack = append(bracketStack, ')')
			} else if char == '{' {
				bracketStack = append(bracketStack, '}')
			} else if char == '[' {
				bracketStack = append(bracketStack, ']')
			} else if char == ')' || char == '}' || char == ']' {
				if len(bracketStack) == 0 {
					bracketStack = append(bracketStack, char)
					break
				}
				if char != bracketStack[len(bracketStack)-1] {
					break
				}
				bracketStack = bracketStack[:len(bracketStack)-1]
			}
		}
		if len(bracketStack) == 0 {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
