package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
		return
	}
	var isLetterPresent [32]int
	options := os.Args[1:]
	for _, v := range options {
		if v[0] == '-' && len(v) >= 2 {
			if v[1] == 'h' {
				fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
				return
			}
			for _, r := range v[1:] {
				if r < 'a' || r > 'z' {
					fmt.Println("Invalid Option")
					return
				}
				isLetterPresent[122-r+6] = 1
			}
		} else {
			fmt.Println("Invalid Option")
			return
		}
	}
	for i, v := range isLetterPresent {
		fmt.Print(v)
		if i != len(isLetterPresent)-1 && (i+1)%8 == 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
