package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		arguments := os.Args[1:]
		for _, word := range arguments {
			if word == "01" || word == "galaxy" || word == "galaxy 01" {
				fmt.Println("Alert!!!")
				break
			}
		}
	}
}
