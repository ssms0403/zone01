package main

import (
	"os"

	"github.com/01-edu/z01"
)

type boolean struct {
	value int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		yes := boolean{value: 0}
		return yes
	} else {
		no := boolean{value: 1}
		return no
	}
}

func even(nbr int) int {
	return nbr % 2
}

func main() {
	lengthOfArg := len(os.Args) - 1
	EvenMsg := "I have an even number of arguments"
	OddMsg := "I have an odd number of arguments"
	if (isEven(lengthOfArg)).value == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
