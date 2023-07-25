package main

import (
	"fmt"
	"os"
)

const SIZE = 2048

func main() {
	if len(os.Args) != 2 {
		return
	}
	brainFuckProgram := []byte(os.Args[1])
	var tabByte [SIZE]byte
	pos := 0
	openBr := 0         // opened brackets
	N := len(brainFuckProgram) // length of the source code
	// iterates through the source code passed in the argument
	for i := 0; i >= 0 && i < N; {
		switch brainFuckProgram[i] {
		case '>':
			// Increment the pointer
			pos++
		case '<':
			// decrement the pointes
			pos--
		case '+':
			// increment the pointed byte
			tabByte[pos]++
		case '-':
			// decrement the pointed byte
			tabByte[pos]--
		case '.':
			// print the pointed byte on std output
			fmt.Printf("%c", rune(tabByte[pos]))
		case '[':
			// go to the matching ']' if the pointed byte is 0 (while start)
			openBr = 0
			if tabByte[pos] == 0 {
				for i < N && (brainFuckProgram[i] != byte(']') || openBr > 1) {
					if brainFuckProgram[i] == byte('[') {
						openBr++
					} else if brainFuckProgram[i] == byte(']') {
						openBr--
					}
					i++
				}
			}
		case ']':
			// go to the matching '[' if the pointed byte is not 0 (while end)
			openBr = 0
			if tabByte[pos] != 0 {
				for i >= 0 && (brainFuckProgram[i] != byte('[') || openBr > 1) {
					if brainFuckProgram[i] == byte(']') {
						openBr++
					} else if brainFuckProgram[i] == byte('[') {
						openBr--
					}
					i--
				}
			}
		}
		i++
	}
}
