package main

import (
	"fmt"
	"os"
)

func main() {
	numArg := len(os.Args)
	if numArg == 2 {
		str := os.Args[1]
		runes := []rune(nil)
		for i := 0; i < len(str); i++ {
			if (i == 0 || i == len(str)-1) && str[i] == ' ' {
				continue
			}
			if str[i] == ' ' && str[i-1] == ' ' {
				continue
			}
			runes = append(runes, rune(str[i]))
		}
		if runes[len(runes)-1] == ' ' {
			fmt.Print(string(runes[:len(runes)-1]))
		} else {
			fmt.Print(string(runes))
		}
	}
	fmt.Println()
}
